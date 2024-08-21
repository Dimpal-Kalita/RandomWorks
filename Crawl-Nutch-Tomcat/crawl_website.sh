#!/bin/bash

# Variables
TOMCAT_HOME="/usr/local/tomcat7"  # Tomcat installation path
CURRENT_DIR=$(pwd)               # Current directory where script is executed
CRAWL_DIR="$CURRENT_DIR/crawl"   # Directory to store crawl data
SEGMENTS_DIR="$CRAWL_DIR/segments"
URLS_FILE="$CURRENT_DIR/urls/seed.txt"        # Make sure to have the URLs listed in this file

# # Inject URLs to be crawled
# echo "Injecting URLs into Nutch..."
# $NUTCH_HOME/bin/nutch inject $CRAWL_DIR/crawl/crawldb $CURRENT_DIR/urls

# # Generate the segments
# echo "Generating segments..."
# $NUTCH_HOME/bin/nutch generate $CRAWL_DIR/crawldb $CRAWL_DIR/crawl/segments 

# # Fetch the content based on the generated segments
# echo "Fetching content..."
# $NUTCH_HOME/bin/nutch fetch $SEGMENTS_DIR/segment-* -dir $CRAWL_DIR

# # Parse the fetched content
# echo "Parsing content..."
# $NUTCH_HOME/bin/nutch parse $SEGMENTS_DIR/segment-* -dir $CRAWL_DIR

# # Update the crawl database
# echo "Updating the crawl database..."
# $NUTCH_HOME/bin/nutch updatedb $CRAWL_DIR/crawldb $SEGMENTS_DIR/segment-* -dir $CRAWL_DIR

# # Check if segments were created
# if [ ! -d "$SEGMENTS_DIR" ]; then
#     echo "Error: Crawl did not generate any segments. Please check your configuration and URL list."
#     exit 1
# fi

mkdir -p $CURRENT_DIR/urls
mkdir -p $CURRENT_DIR/crawl

if [ ! -f $CURRENT_DIR/urls/seed.txt ]; then
    echo "seed.txt not found in $CURRENT_DIR/urls/"
    exit 1
fi

$NUTCH_HOME/bin/nutch inject $CURRENT_DIR/crawl/crawldb $CURRENT_DIR/urls

$NUTCH_HOME/bin/nutch generate $CURRENT_DIR/crawl/crawldb $CURRENT_DIR/crawl/segments

SEGMENT=$(ls $CURRENT_DIR/crawl/segments | tail -n 1)
$NUTCH_HOME/bin/nutch fetch $CURRENT_DIR/crawl/segments/$SEGMENT

$NUTCH_HOME/bin/nutch parse $CURRENT_DIR/crawl/segments/$SEGMENT

$NUTCH_HOME/bin/nutch updatedb $CURRENT_DIR/crawl/crawldb $CURRENT_DIR/crawl/segments/$SEGMENT


# View the crawled data
echo "Viewing the crawled data..."

# List the segments
echo "Listing crawl segments..."
ls -l $SEGMENTS_DIR

# Get the latest segment
LATEST_SEGMENT=$(ls -t $SEGMENTS_DIR | head -1)

if [ -z "$LATEST_SEGMENT" ]; then
    echo "Error: No segments found. The crawling process might have failed."
    exit 1
fi

# Use readseg to display the content of the latest segment
echo "Reading segment data..."
$NUTCH_HOME/bin/nutch readseg -dump $SEGMENTS_DIR/$LATEST_SEGMENT $CURRENT_DIR/dump-output

# Check if the dump-output directory was created
if [ ! -d "$CURRENT_DIR/dump-output" ]; then
    echo "Error: Failed to dump segment data. Please check if the segment path is correct."
    exit 1
fi

# Display fetched content
echo "Displaying fetched content..."
cat $CURRENT_DIR/dump-output/content/*

# Build the WAR file and deploy on Tomcat
echo "Building the Nutch WAR file..."
ant war
WAR_FILE="$NUTCH_HOME/build/nutch-0.9.war"
cp $WAR_FILE $TOMCAT_HOME/webapps/

# Restart Tomcat to deploy the WAR file
echo "Restarting Tomcat to deploy the Nutch web application..."
$TOMCAT_HOME/bin/shutdown.sh
$TOMCAT_HOME/bin/startup.sh

echo "Crawling process completed and Nutch web application deployed!"
echo "You can access the web application at http://localhost:8080/nutch-0.9"

