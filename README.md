<h1 align="center"> ✨ App Versions ✨ </h1>
<p align="center">A web application that shows you the latest version of apps in the respective app store</p>

<p align="center">
  <img src="https://user-images.githubusercontent.com/10229883/216540872-81582dce-d314-41ba-b563-de6d10ce017c.png" width=320 /><img src="https://user-images.githubusercontent.com/10229883/216540882-2c2a52e8-8fe4-418d-88f9-8e99ac009106.png" width=320 /><img src="https://user-images.githubusercontent.com/10229883/216540887-fb920938-28d8-4408-b363-f4f0fd183264.png" width=320 />
</p>

## Examples

| Name | Pretty | Table | JSON |
| - | - | - | - | 
| Android (single app)  - ioki Hamburg | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg&format=table) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg&format=json) | 
| Android (multiple apps) - ioki Hamburg, ioki Wittlich | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&format=table) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&format=json) | 
| iOS (single app) - ioki Hamburg | [Link](https://appversions.vercel.app/?ios=ioki-hamburg/id1400408720) | [Link](https://appversions.vercel.app/?ios=ioki-hamburg/id1400408720&format=table) | [Link](https://appversions.vercel.app/?ios=ioki-hamburg/id1400408720&format=json) | 
| iOS (multiple apps) - ioki Hamburg, ioki Wittlich | [Link](https://appversions.vercel.app/?ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496) | [Link](https://appversions.vercel.app/?ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496&format=table) | [Link](https://appversions.vercel.app/?ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496&format=json) | 
| Android & iOS (multiple apps) - ioki Hamburg, ioki Wittlich | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496&format=table) | [Link](https://appversions.vercel.app/?android=com.ioki.hamburg,com.ioki.wittlich&ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496&format=json) | 


## Formats
The following formats are possible and can be changed with the `format` query parameter:

| Name | Description |
| - | - | 
| pretty (default) | Shows the results in a card-like view |
| table | Shows the result in a table. Each row shows a single app |
| json | Shows the result as a json string. This can also be used as an API for custom integrations |
