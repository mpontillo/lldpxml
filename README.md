# lldpxml

This project aims to digest LLDP XML (as provided from `lldpd`) and provide it
in a more usable format.

## Dependencies

This project uses [goxmlstruct](https://github.com/twpayne/go-xmlstruct) to
generate the XML schema. It must be installed and in the `$PATH` before
regenerating the schema from the test data (see `gen.sh`).
