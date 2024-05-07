# CLI-golang
a comand line interface (CLI) reminder created in golang

The app uses several libraries to handle date and time parsing, and notification.

What the code does:

Initialization: The program checks if the user provided the required arguments. If not, it prints usage instructions and exits.

Parsing Time: Uses the "when" package to parse the first argument and validates that it is a future time. If the time cannot be parsed or is in the past, the program prints an error message and exits.

Handling Reminders: If the program is running in a process that was forked to handle the reminder, it will sleep for the difference between the current time and the provided reminder time, then show the reminder notification using the "beeep" package. The beeep.Alert function is used to display a notification at the specified time, with a custom title, message, and icon.
