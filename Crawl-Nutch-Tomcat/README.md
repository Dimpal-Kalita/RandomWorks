### Steps to Run the Script:

1. **Update Paths:**
   - Replace `$TOMCAT_HOME` and `$NUTCH_HOME` with the actual paths to your Tomcat 7 and Nutch 0.9 installations.

2. **Add URLs:**
   - Update the `URLS_FILE` variable with your seed URLs. Replace `"http://example.com"` with the URLs you want to crawl.

3. **Make the Script Executable:**
   ```bash
   chmod +x crawl_website.sh
   ```

4. **Run the Script:**
   ```bash
   ./crawl_website.sh
   ```
5. **Change the depth and topN values:**
   - You can change the depth and topN values in the script to control the crawl depth and the number of URLs to fetch.
   - The default values are `2` for depth and `5` for topN.
   - It is avaiable `generate` line of the script.
6. **DownGrade Java**
   - [Link](https://stackoverflow.com/questions/9219323/downgrade-java-version)
   ```
   sudo apt-get update 
   sudo apt-get install openjdk-8-jdk
   sudo update-alternatives --config javac
   ```
   - GO to zshrc/ bashrc and change java home address
### Notes:
- Go to "$NUTCH_HOME/conf/nutch-site.xml" to configure Nutch properties.
```xml
<property>
 <name>plugin.includes</name>
  <value>protocol-http|protocol-httpclient|urlfilter-regex|parse-(html)|index-(basic|anchor)|indexer-solr|query-(basic|site|url)|response-(json|xml)|summary-basic|scoring-opic|urlnormalizer-(pass|regex|basic)</value>
</property>
```
- The script configures Nutch, injects URLs, performs the crawl, builds the Nutch WAR file, and deploys it on Tomcat 7.
- Make sure you have Apache Ant installed to build the WAR file.
- You can access the Nutch web application at `http://localhost:8080/nutch-0.9` after the script completes.
1. **Stop the Site:**
    ```bash
    $CATALINA_HOME/bin/shutdown.sh
    ```
