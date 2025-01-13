### Chronofy
A project to help engineers find the root cause of an error by looking through all the possible user traces in the different sources like database, gcp logs, sentry


### Project structure
chronofy
├── cmd/                     # Main application entry point.
│   └── main.go              # Starts the application.
├── internal/                # Application-specific code.
│   ├── providers/           # Data source providers.
│   │   ├── gcp.go           # GCP Logs provider.
│   │   ├── database.go      # Database provider.
│   │   └── sentry.go        # Sentry provider.
│   ├── services/            # Core business logic.
│   │   ├── fetcher.go       # Data fetching logic.
│   │   └── normalizer.go    # Data normalization logic.
│   └── models/              # Data structures.
│   |   └── data.go          # Definitions for normalized data structures.
|   |___handlers             # API handlers
├── pkg/                     # Shared code (reusable outside the app).
│   └── utils/               # Utility functions.
│       └── logger.go        # Logging utility.
├── configs/                 # Configuration files.
│   └── config.yaml          # Example config file.
├── tests/                   # Test files.
│   └── fetcher_test.go      # Unit tests for fetcher logic.
└── go.mod                   # Go module file.


### Run the project

- To run it on production `ENV=production go run cmd/main.go`
- To run it locally `go run cmd/main.go`