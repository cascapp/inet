#!/usr/bin/env python3

if __name__ == "__main__":
    file = open('CURRENT_VERSION',mode='r')
    currentVersion = file.readline()
    file.close()

    splits = currentVersion.split(".")
    major = int(splits[0])
    minor = int(splits[1])
    patch = int(splits[2]) + 1

    file = open('CURRENT_VERSION',mode='w')
    file.write(f"{major}.{minor}.{patch}")
    file.close()

