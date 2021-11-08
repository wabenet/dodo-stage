package spec

include: [...#Include] | *[]

#Include: {
  file: string
}

stages: [string]: #Stage

#Stage: {
  name?:     string
  type:      string
  box:       #Box
  resources: #Resources
  ...
}

#Box: {
  user:          string
  name:          string
  version:       string
  access_token?: string
}

#Resources: {
  cpu:     string
  memory:  string
  volumes: #Volumes | [...#Volume] | *[]
  usb:     #USBFilters |[...#USBFilter] | *[]
}

#Volumes: [string]: #Volume

#Volume: {
  size: string
}

#USBFilters: [string]: #USBFilter

#USBFilter: {
  name:      string
  vendorid:  string
  productid: string
}
