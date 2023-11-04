<h1 align="center"> ✨ AppVers.io ✨ </h1>
<p align="center">A web application that shows you the latest version of apps in the respective app store</p>

<p align="center">
  <img src="https://user-images.githubusercontent.com/10229883/216540872-81582dce-d314-41ba-b563-de6d10ce017c.png" width=320 /><img src="https://user-images.githubusercontent.com/10229883/216540882-2c2a52e8-8fe4-418d-88f9-8e99ac009106.png" width=320 /><img src="https://user-images.githubusercontent.com/10229883/216540887-fb920938-28d8-4408-b363-f4f0fd183264.png" width=320 />
</p>

## Examples

| Name | Pretty | Table | JSON |
| - | - | - | - | 
| Android (single app)  - ioki Hamburg | [Link](https://appvers.io/?android=com.ioki.hamburg) | [Link](https://appvers.io/?android=com.ioki.hamburg&format=table) | [Link](https://appvers.io/?android=com.ioki.hamburg&format=json) | 
| Android (multiple apps) - ioki Hamburg, ioki Wittlich | [Link](https://appvers.io/?android=com.ioki.hamburg,com.ioki.wittlich) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&format=table) | [Link](https://appvers.io/?android=com.ioki.hamburg,com.ioki.wittlich&format=json) | 
| iOS (single app) - ioki Hamburg | [Link](https://appvers.io/?ios=1400408720) | [Link](https://appvers.io/?ios=1400408720&format=table) | [Link](https://appvers.io/?ios=1400408720&format=json) | 
| iOS (multiple apps) - ioki Hamburg, ioki Wittlich | [Link](https://appvers.io/?ios=1400408720,1377071496) | [Link](https://appversions.vercel.app/?ios=1400408720,1377071496&format=table) | [Link](https://appvers.io/?ios=1400408720,1377071496&format=json) | 
| Android & iOS (multiple apps) - ioki Hamburg, ioki Wittlich | [Link](https://appvers.io/?android=com.ioki.hamburg,com.ioki.wittlich&ios=1400408720,1377071496) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&ios=1400408720,1377071496&format=table) | [Link](https://appvers.io/?android=com.ioki.hamburg,com.ioki.wittlich&ios=1400408720,1377071496&format=json) | 
| Android (with Developer ID) - [ioki](https://play.google.com/store/apps/dev?id=8505861851834820244) | [Link](https://appvers.io/?android=did:8505861851834820244) | [Link](https://appvers.io/?android=did:8505861851834820244&format=table) | [Link](https://appvers.io/?android=did:8505861851834820244&format=json) | 
| iOS (with Developer ID) - [ioki](https://apps.apple.com/de/developer/1489448276) | [Link](https://appvers.io/?ios=did:1489448276) | [Link](https://appvers.io/?ios=did:1489448276&format=table) | [Link](https://appvers.io/?ios=did:1489448276&format=json) | 
| Android (with Developer ID) & iOS (with Developer ID) - ioki | [Link](https://appvers.io/?ios=did:1489448276&android=did:8505861851834820244) | [Link](https://appversions.vercel.app/?ios=did:1489448276&android=did:8505861851834820244&format=table) | [Link](https://appvers.io/?ios=did:1489448276&android=did:8505861851834820244&format=json) | 
| Android (with Developer ID + multiple apps) & iOS (with Developer ID + multiple apps) - ioki,  | [Link](https://appvers.io/?ios=did:1489448276,1400408720,1503683596&android=did:8505861851834820244,de.telekom.mobilitysolutions.shuttle,com.ioki.hamburg) | [Link](https://appvers.io/?ios=did:1489448276,1400408720,1503683596&android=did:8505861851834820244,de.telekom.mobilitysolutions.shuttle,com.ioki.hamburg&format=table) | [Link](https://appvers.io/?ios=did:1489448276,1400408720,1503683596&android=did:8505861851834820244,de.telekom.mobilitysolutions.shuttle,com.ioki.hamburg&format=json) | 


## Formats
The following formats are possible and can be changed with the `format` query parameter:

| Name | Description |
| - | - | 
| pretty (default) | Shows the results in a card-like view |
| table | Shows the result in a table. Each row shows a single app |
| json | Shows the result as a json string. This can also be used as an API for custom integrations |
