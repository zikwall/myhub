#!/bin/bash

systemctl stop myhub.service
systemctl start myhub.service
systemctl status myhub.service