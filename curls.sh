#!/bin/bash

curl -H "Content-Type: application/json" -d '{"applicationname":"New App Via Curl", "businessunit": "new business unit"}' http://localhost:8080/applications
