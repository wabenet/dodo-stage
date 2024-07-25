package network

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

const (
	Type = "configure-network"

	dodoBegin = "##_DODO_BEGIN_##"
	dodoEnd   = "##_DODO_END_##"

	interfacesFile = "/etc/network/interfaces"
	networkScripts = "/etc/sysconfig/network-scripts"

	networkManagerConfig = `
DEVICE={{ .Device }}
ONBOOT=yes
BOOTPROTO=dhcp
`

	networkToolsConfig = `
auto {{ .Device }}
iface {{ .Device }} inet dhcp
post-up ip route del default dev $IFACE || true
`
)

type Action struct {
	Device string `mapstructure:"device"`
}

func (a *Action) Type() string {
	return Type
}

func (a *Action) Execute() error {
	log.Printf("configure host network...")

	if _, err := exec.LookPath("nmcli"); err == nil {
		if err := a.configureNetworkManager(); err != nil {
			return err
		}
	} else {
		if err := a.configureNetTools(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Action) configureNetworkManager() error {
	var buffer bytes.Buffer
	templ, err := template.New("network").Parse(networkManagerConfig)
	if err != nil {
		return err
	}
	if err := templ.Execute(&buffer, a); err != nil {
		return err
	}
	configFile := filepath.Join(networkScripts, fmt.Sprintf("ifcfg-%s", a.Device))
	if err := ioutil.WriteFile(configFile, buffer.Bytes(), 0644); err != nil {
		return err
	}

	if nmcli, err := exec.LookPath("nmcli"); err == nil {
		exec.Command(nmcli, "d", "disconnect", a.Device).Run()
	} else {
		exec.Command("/sbin/ifdown", a.Device).Run()
	}

	if systemctl, err := exec.LookPath("systemctl"); err == nil {
		if err := exec.Command(systemctl, "restart", "NetworkManager").Run(); err != nil {
			return fmt.Errorf("could not restart NetworkManager: %q", err)
		}
	} else {
		if err := exec.Command("/etc/init.d/NetworkManager", "restart").Run(); err != nil {
			return fmt.Errorf("could not restart NetworkManager: %q", err)
		}
	}
	return nil
}

func (a *Action) configureNetTools() error {
	// Do not error check because the device might not exist
	exec.Command("/sbin/ifdown", a.Device).Run()
	exec.Command("/sbin/ip", "addr", "flush", "dev", a.Device).Run()

	templ, err := template.New("network").Parse(networkToolsConfig)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	err = templ.Execute(&buffer, a)
	if err != nil {
		return err
	}

	if err := replaceBlockInFile(interfacesFile, dodoBegin, dodoEnd, buffer.String()); err != nil {
		return err
	}

	if err := exec.Command("/sbin/ifup", a.Device).Run(); err != nil {
		return err
	}

	return nil
}

func readSurroundingLines(path string, beginMarker string, endMarker string) ([]string, []string, error) {
	var preLines, postLines []string

	file, err := os.Open(path)
	if err != nil {
		return preLines, postLines, fmt.Errorf("could not open file: %q", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == beginMarker {
			break
		}
		preLines = append(preLines, line)
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == endMarker {
			break
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		postLines = append(postLines, line)
	}

	return preLines, postLines, scanner.Err()
}

func replaceBlockInFile(path string, beginMarker string, endMarker string, contents string) error {
	preLines, postLines, err := readSurroundingLines(path, beginMarker, endMarker)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not open file: %q", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range preLines {
		fmt.Fprintln(writer, line)
	}
	fmt.Fprintln(writer, beginMarker)
	fmt.Fprintln(writer, contents)
	fmt.Fprintln(writer, endMarker)
	for _, line := range postLines {
		fmt.Fprintln(writer, line)
	}
	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
