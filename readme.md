# 🚀 cosmos-linux: Let's Journey through the Cosmos!
Welcome to cosmos-linux, the splendid combination of Cosmos SDK and Celestia DA Layer! Dive into the thrill of creating a Linux VM, execute your favorite commands, and bask in the glow of your machine's output.

## Why the Excitement? 🌌
Because why not? 🤷‍♂️ While running a Linux VM on a rollup might not be the speedy Gonzales of the tech world, it perfectly illustrates the wondrous capabilities of the Cosmos SDK. Plus, it's an exhilarating tutorial for creating rollups and marrying them to Celestia.

## Glittering Features ✨
- [x] Create multiple Linux VMs
- [x] Run commands in the VM
- [x] Read the logs of the VM 
- [x] Caching machines state for querying (on nodes)
- [ ] Deterministic machines
  - [x] Disable network
  - [ ] Change Linux kernel entropy pool for random seeding
  - [ ] Change Date/Time behavior

## Treasure Map 🗺️
Keen to unravel the sorcery behind this project? Your treasure map points to these files:
```
x/cosmoslinux/keeper/*
x/cosmoslinux/utils/docker.go
```
 
## First Steps to the Stars! ⭐️
Ready to blast off? Make sure your spaceship (laptop) is fueled up with Golang and Docker.

Command your local Celestia node to life:
```
docker run \
    -p 36657:26657 -p 36658:26658 -p 36659:26659 -p 9090:9090 \
    ghcr.io/rollkit/local-celestia-devnet:v0.11.0-rc8
```

On another console, ignite the `./init-local.sh` script and soar through the cosmos!

Need some playful trials? Here you go:
```
# create a new machine, you should check the machine id in the logs
cosmos-linuxd tx cosmoslinux create-machine --from linux-key

# run the `ps` command
cosmos-linuxd tx cosmoslinux run-command ${MACHINE_ID} ps --from linux-key

# check the logs of the machine
cosmos-linuxd q cosmoslinux output ${MACHINE_ID}
```
