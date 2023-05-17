# servcli

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

# Specifications

The list of servers will be described in a file in yaml format with the following scheme:

```yaml
servers:
  ITEM1:
    name: Small description for ITEM1
    group1:
      - name: Server1
        isaws: true
      - name: Server2
        isaws: false
      - name: Server3
        isaws: true
    group2:
      - name: Server1
        isaws: false
      - name: Server2
        isaws: true
      - name: Server3
        isaws: true
  ITEM2:
    name: Small description for ITEM2
    group1:
      - name: Server1
        isaws: true
      - name: Server2
        isaws: false
      - name: Server3
        isaws: true
    group2:
      - name: Server1
        isaws: false
      - name: Server2
        isaws: true
      - name: Server3
        isaws: true
```
