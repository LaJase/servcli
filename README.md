# servcli

## Purpose

The "Command line server manager" project is a command line application designed to facilitate server management. Its
main objective is to allow users to select a specific server from a list based on various customizable criteria.

This application provides a user-friendly command line interface that allows users to view and browse a list of
available servers. Users will be able to specify criteria such as geographic location, availability, processing
capacity, current load, or any other attribute relevant to their specific environment.

Through a combination of intuitive commands and advanced filters, users will be able to quickly narrow down the list of
available servers to their specific needs. Once the server is selected, actions such as SSH connection, remote command
execution or resource management can be performed easily.

The "Command line server manager" aims to improve the efficiency and ease of server management, providing administrators
and users with a powerful tool to access and manipulate servers more quickly and accurately.

## Build instructions

If you're new to Go, check out ["How to Build and Install Go Programs"](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)

## Running the Program

navigate to the root directory of the project and run either `go run .` to run the program or `go build` to build the
program binary. If you built the binary, you can run it with `./project-management` or even add it to your PATH so you
can just run `project-management` from anywhere.

## Configuration

For the moment, the configuration is only possible through a yaml file. The program will take by default the one named
`./config/servcli-config.yaml`. But it is possible to provide your own file with the `--config` option.

The list of servers will be described in a file in yaml format with the following scheme:

```yaml
server_list:
  Entity1:
    description: Small description for Entity1
    elements:
      group1:
        - name: e1g1serv1
          isaws: false
        - name: e1g1serv2
          isaws: true
        - name: e1g1serv3
          isaws: false
      group2:
        - name: e1g2serv1
          isaws: true
        - name: e1g2serv2
          isaws: false
        - name: e1g2serv3
          isaws: true
ssh_command: ssh -T git@github.com %s
```

## Collaboration
You're welcome to write features and report issues for this project.
