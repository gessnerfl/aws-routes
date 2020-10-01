# AWS Routes for Mac OS


# Development Details

## Testing

Mocking:
Tests are co-located in the package next to the implementation. We use gomock
(<https://github.com/golang/mock)> for mocking. To generate mocks you need to 
use the package options to create the mocks in the same package:

```bash
mockgen -source=<source_file> -destination=mocks/<source_file_name>_mocks.go -package=mocks -self_package=github.com/gessnerfl/awsroutes/<source_package>
```