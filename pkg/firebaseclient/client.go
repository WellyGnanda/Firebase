package firebaseclient

import (
	"context"
	"encoding/json"

	"go-tutorial-2020/internal/config"
	"go-tutorial-2020/pkg/errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	shadredClient = &firestore.Client{}
	credentials   = map[string]string{
		"type":                        "service_account",
		"project_id":                  "testfirebase-2f50f",
		"private_key_id":              "c8026fb3ef44ce2d10c748bb1ba76dc4f41d3010",
		"private_key":                 "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCvAlfSmMuXgR+e\nbsxVYckozqUQPIIp0wi2AstzB1SrjIuGIuifSSnuToo3jPOJcNj4+u3rlP4sXcT7\nSwo668TPituWk5xIUZog1GWCl+dZ7KZ4qBud/b593lUOypP9TRbtsQPgC4hq6YqY\nLGxXyFjmaXbAcEOQ30eWa0Ay2yWG8KhiKeRqVVolSQfuHVNyjXuTJzl8vHHox7dW\nGpzQb3n6dXq0sovbAZCyfR0MLStZqC9vdqLppqmObnRn+i+IhpjQl/jO2Tlkp5vw\nueBTKl4Z1rdtBI5YHNl2KNVrm3AlyDqI5bVswaqaJOtwnLz1pohS8pWdeAV1G0Dj\nENPb6o2tAgMBAAECggEAFdZ3Qnhd+Tz7xs+BEEtoKD8m4jCGtnTWoq2VGPiqdFCQ\njW+YMt4UjR/AR/++2ODbti/Lleis0bjurkO2FlWapKIpVe/74ZtLHfsa4pGVZQdu\nW2JwtcV2qmqelv6oukQPDyBWQTP3NQ4IxQXQDCEcFL5GuusXR0HRQ0AFTgNB+sUi\ngO9Yyig+WLoR1w+O/oSsDcHzhvp4MhR4sUGjmIa/HMcWe8+HzgbYiUX0AV79oFmz\n1Q4PqylIMjdAR4xPZEBcTWaqN/TiItQJklCqmW6dzpTZi2QR1ZeDihV7+U6SjnJQ\n7DFJpFRmd1HlvyZ7RtRQnXaLk808zGbhj9QG40tqyQKBgQDdcX0TRk532y059fPm\nIlWLbUkoGMubBYCp6B5RYclKtDWWtCk+QuV2cKBxOLl+Gt0qIWHXp3PqFpTu5bbd\nJkB4wgFf5AgPlO4/eppcNsvB0s9F/nQmLXgtDABbNVvzKWTvBXxWsbwuBvefnhCO\nf6FXVaaMM3714H/cNkCWNlDYdQKBgQDKUddIQQmJcMex+qdY8ojjntnnQb2isaOc\nd9wKQEzOXOZgDJaOkpZwOzHJgBKUBtV4f1GSiSi+J0gQxeK32Pb2X+8VXUNP5v5V\n+mqGMoK+dKqDCnrszGntNDpFrIXJYkGtlwRrGZdDH9UF+xBl7DsZcJYYEbk+XFLk\nriQfRYB5WQKBgQCLtgJ3mq//JqVOIEMVOyxFn1m8log+8iXPDMe0CMH7A9+biWdM\nBODI7R4M0QEW8tP+tLkKWnfjhQPKBdxtgqjCh4Ref3wmeIwoOK4S5+9+BgcH3hZh\nz+Y2ZZAD+5JbxA4OT6O2/sP/Nh4c8pj3jsa4Vy2Q3xyG/HEu+nudSf+P0QKBgQCv\nW14P5yb/9DtxfMI9awHQ4CcXtLhL4lHf1VdnnzGzD3wxtddsvYscvYG6l4ICwSWX\nKismqjEhF2Tz/MAz/x6WjrHnv40PHTRGiyR3KiJ+NxpvN88xnT8WdFUpfI387Wfl\nsGYI+gZMDLQTWfdtj+Hte9LsC7iWX2kNgg4W+KORCQKBgD9tAeJsOZOSGV1CYcuR\nYXQyVUQ9tMIBMtRVVnFPHWnaUU1T79gSTuuPHyRdsz/1Utiax5dyNWPeLWw1XZLy\nW8+XyrEfezAmkdKdtJeg2SJKoJGPB74sbzNs2UhsrcCLeKFl1IDlO/hdnk7wos3X\nEtw/N7EGlhaoZpYLfzOxhcSs\n-----END PRIVATE KEY-----\n",
		"client_email":                "firebase-adminsdk-xw35u@testfirebase-2f50f.iam.gserviceaccount.com",
		"client_id":                   "107867893892765728648",
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-xw35u%40testfirebase-2f50f.iam.gserviceaccount.com",
	}
)

type Client struct {
	Client *firestore.Client
}

func NewClient(cfg *config.Config) (*Client, error) {
	var c Client
	cb, err := json.Marshal(credentials)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to marshal credentials!")
	}

	option := option.WithCredentialsJSON(cb)
	c.Client, err = firestore.NewClient(context.Background(), cfg.Firebase.ProjectID, option)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firebase client!")

	}
	return &c, err
}
