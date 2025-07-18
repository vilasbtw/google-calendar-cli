# Google Calendar CLI

A CLI written in Go to fetch your Google Calendar events. You can view what's scheduled for today or for the week.

## Technologies

- [Go](https://golang.org/)  
- [Google Calendar API](https://developers.google.com/calendar/api)  
- [Cobra CLI](https://github.com/spf13/cobra)

## Requirements

- Go  
- A valid `credentials.json` from Google Cloud with access to the Calendar API

## Features

- View today's events: `calendar events today`  
- View this week's events: `calendar events week`  
- Add a calendar to your account: `calendar agenda [agenda_id]`

## Getting Started

1. Clone this repo:

```bash
git clone https://github.com/vilasbtw/google-calendar-cli.git  
```

2. Add your `credentials.json` file from the Google Calendar API to the project root.

3. Run:

```bash
go run main.go events today  
```

Or, list this week's events:

```bash
go run main.go events week  
```

4. To add a specific agenda:

```bash
go run main.go agenda your-agenda-id@group.calendar.google.com  
```

## Command Overview

- `calendar events` — manage events
  - `today` — shows today's events
  - `week` — shows this week's events
- `calendar agenda [agenda_id]` — adds an agenda to your calendar list

## Output

```text
Meeting with team | confirmed | at 2025-07-18T10:00:00-03:00  
Work on CLI app   | confirmed | at 2025-07-18T14:00:00-03:00  
```

