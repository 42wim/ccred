# ccred

Contributor credit generator.

# Features
Displays the contributors of a specific github project between specific commits/tags.
Primary usecase is to get contributors between releases and add them to the release notes/changelog.
Uses `GITHUB_TOKEN` if set.

# Binaries
Binaries can be found [here](https://github.com/42wim/ccred/releases/)

# Usage
```
Usage of ./ccred: ./ccred [commit/tag range options] [owner]/[repo]

Retrieves contributors from a github repository between 2 commits or tags.

Examples:
        ./ccred 42wim/matterbridge
        ./ccred --since v1.11.0 --until v1.12.0 42wim/matterbridge
        ./ccred --since 6f131250f1f48fb3898ee4c6717d9299a215ff67 --until v1.12.0 42wim/matterbridge
        ./ccred --since 6f131250 --until v1.12.0 42wim/matterbridge

Options for commit/tag range:

  -since string
        commit/tag to start from (if not specified takes latest release tag)
  -until string
        commit/tag to end on (default "master")
```

# Example
If no flags specified looks for contributors between last github release and current master.

```
$ ./ccred 42wim/matterbridge
@42wim <Wim>, @DeclanHoare <Declan Hoare>, @KrzysztofMadejski <Krzysiek Madejski>, @AJolly <AJolly>
```

Get golang contributors between go1.11.4 and go1.11.5

```
 $ ./ccred -since go1.11.4 -until go1.11.5 golang/go
@rauls5382 <Raul Silvera>, @ushakov <Max Ushakov>, @Neverik <Stepan Shabalin>, @dmitshur <Dmitri Shuralyov>, @vearutop <Viacheslav Poturaev>, @bmkessler <Brian Kessler>, @jayconrod <Jay Conrod>, @daniel-s-ingram <Daniel Ingram>, @martisch <Martin Möhrmann>, @GuilhermeCaruso <GuilhermeCaruso>, @willbeason <Will Beason>, @jblebrun <Jason LeBrun>, @agnivade <Agniva De Sarker>, @mdempsky <Matthew Dempsky>, @bradfitz <Brad Fitzpatrick>, @mark-rushakoff <Mark Rushakoff>, @dr2chase <David Chase>, @griesemer <Robert Griesemer>, @paulzhol <Yuval Pavel Zholkover>, @ALTree <Alberto Donizetti>, @aarzilli <Alessandro Arzilli>, @jordanrh1 <Jordan Rhee>, @tkivisik <tkivisik>, @Gnouc <LE Manh Cuong>, @catatsuy <catatsuy>, @minux <Shenghou Ma>, @FiloSottile <Filippo Valsorda>, @cannona <Aaron Cannon>, @ianlancetaylor <Ian Lance Taylor>, @alexbrainman <Alex Brainman>, @kardianos <Daniel Theophanes>, @mmcloughlin <Michael McLoughlin>, @josharian <Josh Bleecher Snyder>, @kevinburke <Kevin Burke>, @mvdan <Daniel Martí>, @randall77 <Keith Randall>, @bcmills <Bryan C. Mills>, @andybons <Andrew Bonventre>, @Helflym <Clément Chigot>, @0intro <David du Colombier>, @Quasilyte <Iskander Sharipov>, @mknyszek <Michael Anthony Knyszek>, @gbbr <Gabriel Aszalos>, @mikioh <Mikio Hara>, @yagi5 <Hidetatsu Yaginuma>, @ <Elias Naur>, @aclements <Austin Clements>, @thaJeztah <Sebastiaan van Stijn>, @4a6f656c <Joel Sing>, @oiooj <Baokun Lee>, @juliens <Julien Salleyron>, @cherrymui <Cherry Zhang>, @methane <Inada Naoki>, @neelance <Richard Musiol>, @Inconnu08 <Taufiq Rahman>, @mostynb <Mostyn Bramley-Moore>, @osamingo <Osamu TONOMORI>, @tklauser <Tobias Klauser>, @hyangah <Hana (Hyang-Ah) Kim>
```


# Building
Go 1.17+ is required.

```
go install github.com/42wim/ccred
```

You should now have dt binary in the bin directory:

```
$ ls ~/go/bin/
ccred
```
