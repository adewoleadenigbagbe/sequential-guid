# sequential-guid

A time ordered uuid created essentially for database unique keys

New creates a sequential guid that can be used as the primary key to solve the randomness of Globally Unique Identifier (GUID or UUID).By RFC 4122, the standard GUID is 16 bytes, the first 6 bytes is generated from UTC timestamp and the next 10 bytes are generated randomly,the Endianess are the computer architecture are also considered
Follow up article link: https://www.codeproject.com/Articles/388157/GUIDs-as-fast-primary-keys-under-multiple-database.Golang implementation of code written in C#
