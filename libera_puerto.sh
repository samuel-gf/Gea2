#!/bin/bash
PUERTO=$(lsof -i :8080 | tail -n 1 | cut -f2)
kill -s 9 $(PUERTO)
