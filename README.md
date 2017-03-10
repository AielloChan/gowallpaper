# GoWallpaper

GoWallpaper is a golang build program. It use [Unsplash](https://unsplash.com)'s ppicture source to set your windows wallpaper in a simple and clean way.

# How to use

## Download the excutable file

Actually, this is the most simple and quik operation. You just open the `build` folder and download `Dese.exe` to your computer. Final double click it. 

## Go run

If you already have a golang runtime envirment,you can just clone this repository to your local camputer:

```bash
git clone https://github.com/AielloChan/gowallpaper.git
cd gowallpaper
go run main.go
```

*Note: If you use it in this way, the pictures, download from unsplash, will be stored at system temp directory, not the current folder.*

# Change provider

we provide many picture source, you can see them below:

- [bing每日一图](https://bing.com)
- [百度图片](https://images.baidu.com)
- [Unsplash](https://unsplash.com)
- [爱壁纸](http://aibizhi.com)
- [NationalGeographic国家地理](http://www.nationalgeographic.com/)

You can just add some flag when excute the program to change source:
```bash
// use Unsplash picture (default)
Dese.exe -provider unsplash

// use baidu picture
Dese.exe -provider baidu

// use aibizhi picture
Dese.exe -provider aibizhi

// use bing picture
Dese.exe -provider bing

// use nationalgeographic picture
Dese.exe -provider nationalgeographic
```

# Files

Download pictures will automatically stored at `./pics` folder.