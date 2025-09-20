#set page(paper: "a4", numbering: "1", margin: (inside: 3cm))
#set text(12pt)
#show link: it => underline(text(fill: blue)[#it])

#let proj = "Mnemosyne"
#let version = "v0.1.0"

// Begin Doc

#align(center)[
  #set text(32pt)
  Mnemosyne Specification #version
]
#outline()

= Terms
/ #proj: Project Name
/ Syne: CLI Name
/ Atlas: Name of a filesystem
/ Mnemo: Name of the backend

= Overview
#proj is a file management server with suports for:
- multi-user authentication,
- multiple #proj servers on the client side,
- easy file sharing and management,
- comprehensive tagging system,
- and a simple installation process.

#pagebreak()

= Git Conventions

Use git *tags* / github *releases* for new releases and versions.\
Use a seperate git *branch* for each feature.\
Only merge to the *main* branch.

== Git Commit Header

#let gitmoji(moji, desc) = {
  [
    *:#moji:* - #desc \
  ]
}

#block(inset: (left: 2em))[
  #gitmoji("art", "Improve code structure / formatting / commenting")
  #gitmoji("fire", "Remove code or files")
  #gitmoji("bug", "Fix a bug")
  #gitmoji("ambulance", "Critical hotfix")
  #gitmoji("sparkles", "New feature")
  #gitmoji("memo", "Updated docs")
  #gitmoji("test_tube", "Modify tests")
  #gitmoji("adhesive_bandage", "Non-critical fix")
  #gitmoji("construction", "Work in progress.")
  #gitmoji("recycle", "Refactor code/files")
  #gitmoji("wrench", "Add or update configuration files")
  #gitmoji("beers", "Write code while drunk or otherwise")
  #gitmoji("clown_face", "Mock things")
  #gitmoji("see_no_evil", "Add or update .gitignore")
  #gitmoji("wastebasket", "Deprecate code that needs to be cleaned up")
  #gitmoji("coffin", "Remove dead code")
  #gitmoji("twisted_rightwards_arrow", "Merge fixes")
]

#pagebreak()

= Features

#let cols(cols, body) = context {
  let size = measure(body)
  block(height: size.height / cols + 1em, columns(cols, body))
}

#cols(2)[
  - Mnemo
    - Database
    - File operations
      - Upload file
      - Download file
      - Delete file
      - Query file
      - Rename / move file
    - Authentication
      - Registration
      - Login
      - Secure session cookie
    - User permissions
      - Single user
      - Multi user
      - Temporary shared
    - Networking
    - Docker deployment
  - Syne
    - Login
    - Linux
    - MacOS
    - Windows (?)
    - Directory registration
    - File syncing
    - File sharing
  - GUI
    - Website
    - Mobile app (?)
    - Windows
    - Linux
    - MacOS
    - Same functionality as Syne
    - File preview
]

= Versioning
All versions must follow https://semver.org/

#block(inset: (left: 1em))[
  #quote()[
    Given a version number MAJOR.MINOR.PATCH, increment the:
    MAJOR version when you make incompatible API changes
    MINOR version when you add functionality in a backward compatible manner
    PATCH version when you make backward compatible bug fixes
    Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.]
]

Releases occur for every major / minor version. All releases must include:
- Source code required to build / deploy
- Build scripts and instructions
- Prebuilt `syne` binary for Linux, MacOS, Windows(?), BSD(?)
- GUI applications for Linux, MacOS, Windows, BSD(?)
- Prebuilt `mnemo` binary for Linux, MacOS(?), Windows(?), BSD(?)

#pagebreak()

= Mnemo

== Authentication Architecture
#[
  #show math.equation: set text(fill: green)

  + Clients use Syne/GUI to login to their account system wide. They can be logged into multiple accounts at once. This sends a request to the server which includes the username and password. This will eventually be done over SSL, however at the moment this is plain text
  + Server checks the credentials against all known accounts. If the account exists, it sends back a generated session token that is unique to each login, which is implicitly unique to each client. No information about the client/login is stored other than the time of the login, lifetime of the login, and username
  + Client stores this session token in a system-wide database and associates it with the username used to login.
  + Client then registers a directory to a username. This means that they can be logged into several accounts at the same time, and use different accounts for different directories. This sends a request to the server, which includes the name of the directory and the session token that was retrieved earlier
  + Server checks if the session token is valid, and if not returns a $401$ response to the client. Otherwise it continues as below
  + If the directory already exists, it just returns a success code
  + Otherwise it will create a new directory, then returns a success code
  + Now whenever the client sends a request to the server, it includes the session token in the header.
  + If the session token is still valid, it resets the lifetime
  + If the session token is invalid, it sends a $401$ response to prompt the client to relog
  + If a user that isn’t logged tries to access a file/directory, it will check to see if a) the file/directory has been shared and b) the URL
]

== User Permissions
In order for a user to have access to a directory, they must first register that directory with the server. Once done so, the directory on the server is created, and the account that registered the directory becomes the admin. From then on, the client must provide a session ID that is associated with the account to gain access. To access a directory that isn’t owned by the account, the admin of the directory must grant access to the account. This can be either read, write, or admin. Once access is granted to an account, it can then be set up like usual, however the user must be specified.

#[
  #show math.equation: set text(fill: green)

  Access to a directory can obtained in three ways:
  + The account is the creator of the directory. They are the one that initially registered it with the server, and as such have admin privileges.
  + The account is granted access by an admin of a directory. This can be one of three levels below. This is represent internally similarly to how Unix represents file permissions, except any permission level is assumed to include any lower permission levels. These are $0$ for no access, $1$ for read access, $2$ for write access, and $3$ for admin access.
    + Read - They can read, preview, and download any files within the directory or subdirectories that inherit the parent permissions
    + Write - They can do everything that a reader can do, as well as edit, move, rename, delete, or upload any file to the directory or subdirectory that inherit the parent permissions
    + Admin - They can do everything a writer can do, as well as manage access to the directory or subdirectory that inherit the parent permissions.
  + Anyone with access to the directories normally can share a URL or temporary password that links to a single file or directory that can be viewed for a limited time or number of views. These shared
]

=== Networking Specification
==== Response Codes

#{
  let response(num, body) = {
    [#num - #body]
  }

  response(401, "Invalid session token")
}

#pagebreak()

== Syne Specification

=== Global Flags

- help, h : Show help
- guest, g: Execute command as a guest
- server=STRING, s: Manually specify server name or url

=== Commands
All commands act on the server specified by the server flag, if the server is not set, it will look for an initialised directory and use that server.

Relative paths inside of an initialised directory expand to the current folder the user is in.

Absolute paths beginning in “/” in an initialised directory and out expand to the root of the server.

Absolute paths beginning in “\~” expand to the root of the initialised directory. Does not work outside of initialised directories.

Server registry file will be searched for at the `SYNE_REGISTRY` environment variable, or in `~/.syne.json`.

#[
  #show raw: it => text(fill: blue, it)
  - `register <name> <url>`
    - Registers a server with a given name
  - `login`
    - Logs into the server
    - Either uses current initialised dir or server flag
  - `logout [--others,-o]`
    - logs out of the current initialised dir server, or server flag
    - if [others] flag set, logout all others from curr account
  - `logoutall`
    - logs out of all accounts
  - `init [<path>]`
    - initialises current directory as a synced folder
    - Uses path as a sub-dir on the server, defaults to root
  - `sync [<path>]`
    - uploads the specified path in current initialised directory to the server or the whole folder by default
  - `get [<path>]`
    - gets the specified resource, defaults to root
  - `add <from> [<to>]`
    - uploads file / folder from to the server at to
  - `mv <from> <to>`
    - renames file / folder from to to on the server
  - `del <path>`
    - deletes a file / folder on the server at path
  - `info <path>`
    - gets info about a file / folder on the server
  - `share <path> [--password,-p]=STRING [--num-uses, -n]=INT [--alive-time,-t]=INT`
    - Shares a path with an optional password, number of uses and alive time.
]

#pagebreak()

== GUI Spec

