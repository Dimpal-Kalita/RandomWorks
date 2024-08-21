### Steps to Run the Script:

1. **Update Paths:**
   - Replace `/path/to/tomcat7` and `/path/to/nutch-0.9` with the actual paths to your Tomcat 7 and Nutch 0.9 installations.

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

### Notes:
- The script configures Nutch, injects URLs, performs the crawl, builds the Nutch WAR file, and deploys it on Tomcat 7.
- Make sure you have Apache Ant installed to build the WAR file.
- You can access the Nutch web application at `http://localhost:8080/nutch-0.9` after the script completes.
