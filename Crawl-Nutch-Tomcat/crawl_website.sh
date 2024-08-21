h

# Variables (update these paths and URLs as needed)
TOMCAT_HOME="/usr/local/tomcat7"  # Set your actual Tomcat installation path
NUTCH_HOME= "~/nutch-0.9"  # Set your actual Nutch installation path
URLS_FILE="urls/seed.txt"        # Make sure to have the URLs listed in this file

# Navigate to the Nutch home directory
cd $NUTCH_HOME

# Configure Nutch (minimal configuration example)
echo "Configuring Apache Nutch..."
cat > conf/nutch-site.xml <<EOL
<configuration>
  <property>
    <name>http.agent.name</name>
    <value>Nutch</value>
  </property>
  <property>
    <name>http.agent.description</name>
    <value>Apache Nutch</value>
  </property>
  <property>
    <name>http.agent.url</name>
    <value>http://nutch.apache.org/</value>
  </property>
  <property>
    <name>http.agent.email</name>
    <value>your-email@example.com</value>
  </property>
</configuration>
EOL

# Inject URLs to be crawled
echo "Injecting URLs into Nutch..."
mkdir -p urls
echo "http://example.com" > $URLS_FILE  # Add your URLs here
bin/nutch inject $URLS_FILE

# Start crawling
echo "Starting the crawl..."
bin/nutch crawl urls -depth 3 -topN 100

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

