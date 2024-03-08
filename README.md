<!-- LTeX: language=en-US -->
# MEGADUNGEON FACTIONS

This project is a Megadungeon Faction Generator for all kinds of TTRPGs.
It is the result of a [vote by the people of the r/OSR](https://www.reddit.com/r/osr/comments/1b6a3d8/a_software_dev_wants_to_give_back_to_this/)
community.

> This project is always open for contributions, especially for new content
> packs; so go ahead and add the things you'd like to see!

You can find this service hosted on [TODO](TODO).

I also have a [discord community](https://discord.gg/f8nyyK4ngZ), join if you'd
like to partake in conversation about this and other projects of mine!

If you want to show some love for what I do, or want help pay the server, you can...<br>
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/yellow_img.png)](https://www.buymeacoffee.com/winterv)

## Running Locally

If you want to [contribute](#contributing), or just run the project locally for
your own use, follow the instructions below.

> This project is built and maintained on Linux. While I don't think it's
> generally impossible to run on Windows, but the
> [Taskfile](https://taskfile.dev/) is written using Linux commands.

### Tooling Dependencies

- [Golang](https://go.dev/)
- [Task](https://taskfile.dev/)
- [Docker](https://www.docker.com/)
- [gcloud](https://cloud.google.com/sdk/gcloud)

### Running The Webserver

Use the included [Taskfile](https://taskfile.dev/) to run these services.

Use this to run the project locally in a docker container.
```bash
task local-webserver
```

If you'd like to run without docker, but with an automatic rebuild when
`.templ` templates change, run:
```bash
task dev
```

## Deploying in the Cloud
The project was created with the intention of hosting on Google Cloud Run.
Google Cloud Run can be replaced with any serverless platform, but some work
will be required if this is your goal, and the following instructions will
assume Google Cloud services.

To deploy to Google Cloud, first set all instances of the `PROJECT_ID` variable
in the included Taskfile to your project ID.

You may then use `task gcloud-setup` to initialize some `gcloud` settings.
(Make sure you have `gcloud` installed!)

Then, in order to deploy, run `task deploy-webserver`. That should be it!

## Testing
Tests can be run by using the included [Taskfile](https://taskfile.dev/).
```bash
task test
```

## Techstack
- [Go](https://go.dev/)
- [Templ](https://github.com/a-h/templ)
- [HTMX](https://htmx.org/)
- [Google Cloud](https://cloud.google.com/?hl=en)
- [Docker](https://www.docker.com/)

## Contributing
- Before posting a pull request, please use [`go fmt`](https://go.dev/blog/gofmt)
    to format your code.
- Beginners to open source are welcome. If you'd like to contribute, but don't
    understand something, you're welcome to ask using an issue.
- Please post feature requests as one issue per feature.
- Before working on a larger contribution, please open an issue to ask if the
    feature you want to implement would be welcome.
