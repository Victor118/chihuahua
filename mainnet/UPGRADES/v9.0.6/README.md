  GNU nano 8.4                                                                        mainnet/UPGRADES/v9.0.6/README.md                                                                                 
# Chihuahua v9.0.6 Upgrade

The Upgrade is scheduled for block `20523000`. A countdown clock is [here](https://www.mintscan.io/chihuahua/blocks/20523000)

This guide assumes that you use cosmovisor to manage upgrades.

In the #96 Proposal you will find a different version hash, the correct one is the one below: 15f53f0ab90f1c251099b3012e2946de1188091f

## If you are syncing from 0 you need to apply v9.0.6 at height 20523000

```bash
# get the new version
cd chihuahua   
git fetch --all    
git checkout v9.0.6
make install
```

# check the version

```bash
# should be v9.0.6
chihuahuad version
# Should be commit 15f53f0ab90f1c251099b3012e2946de1188091f
chihuahuad version --long | grep commit
```

# Make new directory and copy binary

```bash
mkdir -p $HOME/.chihuahuad/cosmovisor/upgrades/v9.0.6/bin
cp $HOME/go/bin/chihuahuad $HOME/.chihuahuad/cosmovisor/upgrades/v9.0.6/bin
```

# check the version again

```bash
# should be v9.0.6
$HOME/.chihuahuad/cosmovisor/upgrades/v9.0.6/bin/chihuahuad version
