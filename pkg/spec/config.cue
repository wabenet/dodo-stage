package spec

stages: [string]: #Stage

#Stage: {
  name?:      string
  type:       string | *"fixed"
  provision?: #Provision

  if type == "fixed" {
    address:   string
    ca_path:   string
    cert_path: string
    key_path:  string

    ssh_config?: #SSHConfig
  }

  ...
}

#SSHConfig: {
  host:             string
  port:             int
  username:         string
  private_key_file: string
}

#Provision: {
  type: string | *"none"

  if type == "stagehand" {
    stagehand_url: string
    script:        [...string] | *[]
  }

  ...
}
