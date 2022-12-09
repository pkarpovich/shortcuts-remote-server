# Shortcuts Remote Server

Simple server to run a shortcut from iPhone or iPad on a remote MacOs computer

## Example

Run shortcuts on remote computer
```bash
GET http://{{url}}/execute?name={{shortcut_name}}
```

Note: The shortcut with `shortcut_name` must be created in the `Shortcuts` app of the user running the server
