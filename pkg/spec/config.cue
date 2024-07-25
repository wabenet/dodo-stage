package spec

stages: [string]: #Stage

#Stage: {
  name?:      string
  type:       string | *"fixed"
  provision?: #Provision
  ...
}

#Provision: {
  type: string | *"stagehand"

  if type == "fixed" {
    address:   string
    ca_path:   string
    cert_path: string
    key_path:  string
  }

  if type == "stagehand" {
    stagehand_url: string
    script:        [...string] | *[]
  }

  ...
}
