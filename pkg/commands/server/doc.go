/*
Package server contains the CLI command definitions for interacting with an
instance of Synse Server. The commands here correlate to the JSON API core
endpoints provided by Synse Server, namely:

 - /test
 - /version
 - /config
 - /scan/...
 - /read/...
 - /write/...
 - /transaction/...
*/
package server

// SynseActionsCategory defines the category name for Synse Server actions.
const SynseActionsCategory = "Synse Server Actions"
