package spec

stages: [string]: #Stage

#Stage: {
  name?:      string
  type:       string
  provision?: #Provision
  ...
}

#Provision: {
  stagehand_url: string
  script:        [...string] | *[]
}
