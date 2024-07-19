#!/bin/bash

goxmlstruct \
    -char-data-field-name Data \
    -package-name lldpxml \
    testdata/example.xml > \
        xmlstruct.go
