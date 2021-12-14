<p align="center">
  <img width="200" height="250" src="https://cdn.pixabay.com/photo/2019/02/24/18/30/pirate-4018168_1280.png">
</p>

# Commandarrgh in a nuthsell

Commandarrgh is an interface that helps you marshaling data into a command arguments structure. Maybe you have been trying to use others cli frameworks to make this task without any effect because all of them are based on the user input.

# How to use it

Commandarrgh is easy to use, you only need to define the struct in which the data will be marshaled to by using **commandarrgh** tags:

```go
type CommandTransferFilesArgs struct {
    Dst `flag:"-dst" helptext:"file's directory" binding:"required" validators:"stat-path"`
    Src `flag:"-src" helptext:"file's directory" binding:"required" validators:"stat-path"`
    AlsoGetHiddens `flag:"-h" helptext:"include hidden folder's files" binding:"bool" default:"false" `
}
```

Method `Marshal` will fill the struct with the string data as the following example shows:

```go
args := &CommandTransferFilesArgs{}
data := "-dst ./memes/dog_bonks -src /downloaded_memes/dog_bonks -h"

err := commandarrgh.Marshal(data, args)
...
```

A full example of how to use **commandarrgh**:

```go
package main

import github.com/alvarogf97/commandarrgh

const (
    COMMAND_TRANSFER_FILES = "transfer"
)

type CommandTransferFilesArgs struct {
    Dst `flag:"-dst" helptext:"file's directory" binding:"required" validators:"stat-path"`
    Src `flag:"-src" helptext:"file's directory" binding:"required" validators:"stat-path"`
    AlsoGetHiddens `flag:"-h" helptext:"include hidden folder's files" binding:"bool" default:"false" `
}

func handleTransferFilesCommand(args CommandTransferFilesArgs){
    // handle here your command :D
    ...
}

func main(){
    data := "transfer -dst ./memes/dog_bonks -src /downloaded_memes/dog_bonks -h"
    cmd, cmdargs := commandarrgh.SplitCommand(data)

    switch cmd{
        case COMMAND_TRANSFER_FILES:
            args := &CommandTransferFilesArgs{}
            err := commandarrgh.Marshal(cmdargs, args)
            if err != nil {
                panic(err)
            }
            handleTransferFilesCommand(args)
        default:
            panic(fmt.Sprintf("No command handler found for `%s`", cmd))
    }
}
```

# Validators

**Commandarrgh** helps you with data validation and trasnformation by using
**validators** which are functions that gives the original value and returns
the modified one or error in case the validation goes wrong. The following validators
are available:

| Name| Function
|-----------|-------------------------------------------------------------------------------------------------------------------
| path| verifies the given value is a wellformed path, trasnform it to an absolute one and make it compatible with all os 
| stat-path| makes the same function that **path** but it must be exists in the system

If someones need another validatios, please write an issure or a pull request and I will work to make it real
