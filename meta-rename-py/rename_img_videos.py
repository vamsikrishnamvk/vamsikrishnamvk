import os
import re
import subprocess
from datetime import datetime


# Function to extract metadata using ExifTool
def get_metadata(file_path):
    try:
        # Use ExifTool to extract metadata as JSON
        result = subprocess.run(
            ['exiftool', '-j', file_path],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            universal_newlines=True,
        )
        if result.returncode != 0:
            print(f"Error reading metadata for {file_path}: {result.stderr}")
            return None
        import json
        metadata = json.loads(result.stdout)[0]  # Parse the first (and only) file's metadata
        return metadata
    except Exception as e:
        print(f"Error while extracting metadata: {e}")
        return None


# Function to convert GPS to decimal degrees
def convert_gps_to_decimal(gps_data):
    try:
        if not gps_data:
            return None
        # Convert the DMS (degrees, minutes, seconds) format to decimal degrees
        parts = re.split(r"[^\d.]+", gps_data.strip())
        if len(parts) < 4:
            return None
        degrees = float(parts[0])
        minutes = float(parts[1])
        seconds = float(parts[2])
        direction = gps_data.strip()[-1].upper()
        decimal = degrees + (minutes / 60.0) + (seconds / 3600.0)
        if direction in ['S', 'W']:
            decimal *= -1
        return f"{decimal:.6f}"  # Format to 6 decimal places
    except Exception as e:
        print(f"Error converting GPS: {gps_data}. Error: {e}")
        return None


# Function to format metadata into a filename
def format_filename(metadata, fallback_date, original_name, extension):
    # Check for CreateDate
    create_date = metadata.get('CreateDate', None)
    if create_date and create_date != "0000:00:00 00:00:00":
        create_date = re.sub(r'[:]', '-', create_date)  # Replace invalid colon characters
        create_date = datetime.strptime(create_date, "%Y-%m-%d %H-%M-%S")
    else:
        create_date = fallback_date  # Use file modification date as fallback

    # Check for GPS coordinates and convert to decimal degrees
    gps_latitude = convert_gps_to_decimal(metadata.get('GPSLatitude', ''))
    gps_longitude = convert_gps_to_decimal(metadata.get('GPSLongitude', ''))

    if gps_latitude and gps_longitude:
        # Format filename with CreateDate and GPS
        return f"{create_date.strftime('%Y-%m-%d_%H-%M-%S')}_{gps_latitude}_{gps_longitude}{extension}"
    elif create_date:
        # Format filename with CreateDate only
        return f"{create_date.strftime('%Y-%m-%d_%H-%M-%S')}_NoGPS{extension}"
    else:
        # Fallback to original filename
        return f"{original_name}_NoMetadata{extension}"


# Main function to rename files in a directory
def rename_files_in_directory(directory):
    if not os.path.exists(directory):
        print(f"The directory '{directory}' does not exist.")
        return

    # Process each file in the directory
    for root, _, files in os.walk(directory):
        for file in files:
            file_path = os.path.join(root, file)
            # Skip unsupported files
            if not file.lower().endswith(('.mp4', '.mov', '.jpg', '.jpeg', '.png')):
                continue

            # Extract metadata
            metadata = get_metadata(file_path)
            if metadata is None:
                continue

            # Get file modification date as fallback
            file_mod_date = datetime.fromtimestamp(os.path.getmtime(file_path))

            # Get the file extension
            extension = os.path.splitext(file)[1].lower()  # Retain original extension (.mp4, .mov, etc.)

            # Generate new filename
            new_name = format_filename(metadata, file_mod_date, os.path.splitext(file)[0], extension)
            new_path = os.path.join(root, new_name)

            # Rename the file
            try:
                os.rename(file_path, new_path)
                print(f"Renamed: {file_path} -> {new_path}")
            except Exception as e:
                print(f"Failed to rename {file_path}: {e}")


# Run the script
if __name__ == "__main__":
    import sys
    if len(sys.argv) != 2:
        print("Usage: python rename_files.py <directory>")
        sys.exit(1)

    target_directory = sys.argv[1]
    rename_files_in_directory(target_directory)
