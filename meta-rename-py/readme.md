# Video File Renaming Script

This Python script renames video files (e.g., `.mp4` and `.mov`) based on their metadata. The new filenames include the creation date, GPS coordinates (if available), and fallback to the file modification date if metadata is incomplete.

---

## Features

- **Rename Files Dynamically**: Extracts metadata from video files using `exiftool`.
- **File Types Supported**:
  - Video files: `.mp4`, `.mov`
  - Image files: `.jpg`, `.jpeg`, `.png`
- **Smart Metadata Handling**:
  - Includes GPS coordinates in filenames when available.
  - Falls back to file modification date if no metadata is present.
- **Filename Format**:
  - `YYYY-MM-DD_HH-MM-SS_Latitude_Longitude.extension` (if GPS is available).
  - `YYYY-MM-DD_HH-MM-SS_NoGPS.extension` (if GPS is unavailable).
  - `YYYY-MM-DD_HH-MM-SS_NoMetadata.extension` (if no metadata is available).

---

## Documentation

- [Usage Instructions](usage.md)
- [Examples](examples.md)
- [Supported File Types](supportedfiletypes.md)
- [Troubleshooting](troubleshooting.md)
- [Author](author.md)

---

## License

This script is open-source and free to use.
