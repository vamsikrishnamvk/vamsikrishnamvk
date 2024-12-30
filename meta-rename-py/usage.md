# Usage Instructions

### Running the Script
1. Open a terminal or command prompt.
2. Navigate to the folder where `rename_files.py` is saved:
   ```bash
   cd C:\Path\To\Folder
3. Run the script by providing the directory containing the video files as an argument
   ```python
   python rename_files.py "C:\Path\To\Videos"
### Script Behaviour
* The script will scan the specified directory for .mp4, .mov, .jpg, .jpeg, and .png files.
* It will rename the files based on their metadata:
    * If CreateDate and GPS data are available: YYYY-MM-DD_HH-MM-SS_Latitude_Longitude.extension
    * If only CreateDate is available: YYYY-MM-DD_HH-MM-SS_NoGPS.extension
    * If no metadata is available, it will fall back to the file modification date: YYYY-MM-DD_HH-MM-SS_NoMetadata.extension