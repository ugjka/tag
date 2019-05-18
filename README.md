# MP3/MP4/OGG/FLAC metadata parsing library
[![GoDoc](https://godoc.org/github.com/ugjka/tag?status.svg)](https://godoc.org/github.com/ugjka/tag)

This package provides MP3 (ID3v1,2.{2,3,4}) and MP4 (ACC, M4A, ALAC), OGG and FLAC metadata detection, parsing and artwork extraction.

Detect and parse tag metadata from an `io.ReadSeeker` (i.e. an `*os.File`):

```go
m, err := tag.ReadFrom(f)
if err != nil {
	log.Fatal(err)
}
log.Print(m.Format()) // The detected format.
log.Print(m.Title())  // The title of the track (see Metadata interface for more details).
```

Parsed metadata is exported via a single interface (giving a consistent API for all supported metadata formats).

```go
// Metadata is an interface which is used to describe metadata retrieved by this package.
type Metadata interface {
	Format() Format
	FileType() FileType

	Title() string
	Album() string
	Artist() string
	AlbumArtist() string
	Composer() string
	Genre() string
	Year() int

	Track() (int, int) // Number, Total
	Disc() (int, int) // Number, Total

	Picture() *Picture // Artwork
	Lyrics() string
	Comment() string

	Raw() map[string]interface{} // NB: raw tag names are not consistent across formats.
}
```

## Audio Data Checksum (SHA1)

This package also provides a metadata-invariant checksum for audio files: only the audio data is used to
construct the checksum.

[http://godoc.org/github.com/ugjka/tag#Sum](http://godoc.org/github.com/ugjka/tag#Sum)
