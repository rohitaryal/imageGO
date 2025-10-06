# imagego

Re-implementation of lab.google's unofficial imageFX API in golang

## Usage

### Importing as a dependency

```go
// Please AI generate yourself :(
// Or wait until my ATP regenerates
```

### Using as CLI

Download latest binary from release and run the following:

```bash
chmod +x imagego
./imagego --image --prompt "purple cat" --cookie "$GOOGLE_COOKIE"
```

Full usage:

```text
imagego [flags] --cookie [cookie]

  -caption
        Generate description from an image
  -cookie string
        User account cookie
  -count int
        Number of images to generate (default 1)
  -dir string
        Destination directory to save images (default ".")
  -fetch
        Fetch generated images using unique media ID
  -id string
        Unique media generation id
  -image
        Generate image from a prompt
  -model string
        Model to use for generation (default "IMAGEN35")
  -prompt string
        Textual description for the image (default "Purple cat")
  -seed int
        A specific number that serves as the starting point
  -size string
        Aspect ratio of the image (default "LANDSCAPE")
  -verbose
        Extra logs

Available sizes: SQUARE, LANDSCAPE, PORTRAIT, UNSPECIFIED
Available models: IMAGEN3, IMAGEN31, IMAGEN35
```

### Help

<details>
<summary style="font-weight: bold;font-size:15px;">How to extract cookies?</summary>

#### Easy way

1. Install [Cookie Editor](https://github.com/Moustachauve/cookie-editor) extension in your browser.
2. Open [labs.google](https://labs.google/fx/tools/image-fx), make sure you are logged in
3. Click on <kbd>Cookie Editor</kbd> icon from Extensions section.
4. Click on <kbd>Export</kbd> -> <kbd>Header String</kbd>

#### Manual way

1. Open [labs.google](https://labs.google/fx/tools/image-fx), make sure you are logged in
2. Press <kbd>CTRL</kbd> + <kbd>SHIFT</kbd> + <kbd>I</kbd> to open console
3. Click on <kbd>Network</kbd> tab at top of console
4. Press <kbd>CTRL</kbd> + <kbd>L</kbd> to clear network logs
5. Click <kbd>CTRL</kbd> + <kbd>R</kbd> to refresh page
6. Click on `image-fx` which should be at top
7. Goto <kbd>Request Headers</kbd> section and copy all the content of <kbd>Cookie</kbd>

</details>

<details>
<summary style="font-weight: bold;font-size:15px;">ImageFX not available in your country?</summary>

1. Install a free VPN (Windscribe, Proton, etc)
2. Open [labs.google](https://labs.google/fx/tools/image-fx) and login
3. From here follow the "How to extract cookie?" in [HELP](#help) section (above).
4. Once you have obtained this cookie, you don't need VPN anymore.

</details>

<details>
<summary style="font-weight: bold;font-size:15px;">Not able to generate images?</summary>

Create an issue [here](https://github.com/rohitaryal/imageGO/issues). Make sure the pasted logs don't contain cookie or tokens.
</details>

## Contributions

Contribution are welcome but ensure to pass all test cases and follow existing coding standard.

## Disclaimer

Bare minimum effort has been put into this project as this is just re-writing. For better support please use [imagefx-api](https://github.com/rohitaryal/imageFX-api)
This project demonstrates usage of Google's private API but is not affiliated with Google. Use at your own risk.

