# gross
gross is a auto cross compile framework, which uses docker to build firmwares.

### purpose
Building cross toolchain for embeded development has always been annoying in traditional ways(may even cause days of efforts). And sometimes we found it hard to make the building pross automatic. We need a more flexible and unmanned firmware delivery pattern. 

### general ideas
So we create gross, with mechanism like below:

1. Use docker to hold firmware cross toolchain/sdk. This is a more lightweight aproach than the traditional virtual machine, and can invoke synchronized and isolated build task at a time.
2. Integrated to popular git based CVS platforms like github/gitlab. So the building and distributing process can be trigered automatically by source pushing or tagging.
3. The docker builds' output target bin can be discovered by gross. We use file watcher to detect folder change to get the bin files and notify the distributing process the handle the built firmware. 

### architecture overview
TODO