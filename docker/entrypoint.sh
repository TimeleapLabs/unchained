#!/bin/sh

unchained postgres migrate conf.yaml && unchained start conf.yaml --generate
