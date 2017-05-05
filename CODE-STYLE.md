# eWitness Code Style Guide
In order to facilitate readable and easy to maintain code, this document proposes a guide to the coding style that eWitness developers and developers submitting pull requests to eWitness should follow. In addition to vetting for bugs, efficiency, and security, pull requests will also be vetted based on code style. This document will be structured as follows:
1. File Structure
2. General Code Style
3. Function Size and Documentation
4. Struct Documentation

## File Structure
In order to maximize readability and auditability of code, files should follow some conventions. This section closes with an example of a file that fits the correct file structure.

Firstly, the top of every file should contain a multi-line comment header with the following information:
1. The name of the file.
2. The list of primary authors.
3. The purpose of the file.

Secondly, files should contain collections of structs and functions that are related to some single task. Files with mixed purposes are confusing to individuals trying to audit the code. If a funciton you are implimenting does not logically fit in any of the existing files, please create a new file for it. One can tell if a function fits the purpose of a particular file by looking at the file's comment header.

After the file header, the package name and imports should be defined. Imports should each be on its own line, as is typically seen in Go code.

Fourthly, constants should be defined. To make it clear that a value is a constant, its name should be in all capitals. Constants should be preceeded by a code comment that makes it clear that the following lines define constants.

Any structs that belong in the file should appear after the constants. For maximum readability, it is preferred that at most one struct exist in a given file. Having more than one struct per file is acceptable in certain situations, such as when one struct is contained in another, but this should be the exception rather than the rule. For more information about structs, please see the section on Struct Documentation.

Lastly, files should contain any function definitions. Functions should fit the theme of the file, and any functions that primarily operate on the struct defined in a file should be in the same file. For more information about functions, please see the section on Function Size and Documentation.

An example of an eWitness go file can be seen below. This files does not contain a struct.

```
/*
    File: Messaging/messaging.go
    Authors: Kevin Gallagher

    Messagining is a package of funcitons to handle the network
    communication between servers. It provides functionality for sending
    data down the wire and for handling networking communication that
    arrives, calling the appropriate functions. It does not prepare data
    for transfer across the wire by converting to JSON. That must be done
    elsewhere.
*/

package Messaging

import (
    "fmt"
    "net"
    "io/ioutil"
)

/*
    Below are the INT assignments to given functions that will appear
    in the messaging struct. These values should not change.
*/
const ADD_HASH              := 1
.
.
.
/*
  HandleSocket will handle a connection when it comes in.

  Args:   1) A connection to the remote client. (type net.Conn)
          2) A database connection to run queries on. (type string)

  Input:  None
  
  Output: 1) Network output to the client. This can be an acknowledge
             message, a failure message, or an authentication challenge.
          2) Storage output to the database referred to by input 2. This
             can be a hash record, a user record, a fetch request update,
             or more.
          3) An integer to indicate success or failure of the function's
             execution.
  Return: None

*/
func HandleSocket(connection net.Conn, databaseConnection string) {
   defer connection.Close()
   .
   .
   .
}
.
.
.

```

## General Code Style

## Function Size and Documentation

## Struct Documentation 
