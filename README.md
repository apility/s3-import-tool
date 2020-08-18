# Netflex Import Tool

A tool that helps sync data to netflex.

Contact Apility support staff in order to get access keys and bucket name.


## Importing data to netflex.

To use the tool, set the bucketname by using the `-b` flag. Then give it a list of all files that needs to be synced.

```
$ netflex-import import -b <bucketname> File [File2]
``` 

It is also possible to use wildcards such as

```
$ netflex-import import -b <bucketname> Folder/*.json
``` 

to sync all json files