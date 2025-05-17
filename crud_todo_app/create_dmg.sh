#!/bin/bash

# Set variables
APP_NAME="Todo App"
DMG_NAME="${APP_NAME}.dmg"
APP_PATH="./${APP_NAME}.app"
DMG_PATH="./${DMG_NAME}"
VOLUME_NAME="${APP_NAME}"

# Create a temporary directory for the DMG contents
TMP_DIR="./tmp_dmg"
mkdir -p "${TMP_DIR}"

# Ensure the app bundle is properly structured
if [ -d "${APP_PATH}" ]; then
    # Copy the app to the temporary directory
    cp -R "${APP_PATH}" "${TMP_DIR}/"
    
    # Rename the executable to match the app name
    if [ -f "${TMP_DIR}/${APP_NAME}.app/Contents/MacOS/crud_todo_app" ]; then
        mv "${TMP_DIR}/${APP_NAME}.app/Contents/MacOS/crud_todo_app" "${TMP_DIR}/${APP_NAME}.app/Contents/MacOS/${APP_NAME}"
    fi
    
    # Create the DMG
    hdiutil create -volname "${VOLUME_NAME}" -srcfolder "${TMP_DIR}" -ov -format UDZO "${DMG_PATH}"
    
    # Clean up
    rm -rf "${TMP_DIR}"
    
    echo "DMG file created: ${DMG_PATH}"
else
    echo "Error: App bundle not found at ${APP_PATH}"
    exit 1
fi 