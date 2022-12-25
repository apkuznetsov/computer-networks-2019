package webcrawler;

import static webcrawler.Webcrawling.MAX_DEPTH_NUMBER;
import static webcrawler.Webcrawling.MIN_DEPTH_NUMBER;
import static webcrawler.WebcrawlingParsing.parseHtmlCode;
import static webcrawler.WebcrawlingParsing.parseLinksFromHtmlCode;
import static webcrawler.WebcrawlingParsing.parseTitle;

import java.util.List;
import java.util.Queue;
import java.util.Set;

public class WebcrawlingThread extends Thread {

  private final Queue<String> tasks;
  private final Set<String> processedUrls;

  private final List<String> urls;
  private final List<String> urlsTitles;

  private final int depthNumber;

  public WebcrawlingThread(
      final List<String> urls, final List<String> urlsTitles,
      final Queue<String> tasks, final Set<String> processedUrls,
      final int depthNumber) {

    if (depthNumber < MIN_DEPTH_NUMBER
        || depthNumber > MAX_DEPTH_NUMBER) {
      throw new IndexOutOfBoundsException();
    }

    this.tasks = tasks;
    this.processedUrls = processedUrls;

    this.urls = urls;
    this.urlsTitles = urlsTitles;

    this.depthNumber = depthNumber;
  }

  @Override
  public void run() {
    int repeatTimes = depthNumber;
    do {
      final String currUrl = tasks.poll();
      if (currUrl == null
          || processedUrls.contains(currUrl)) {
        return;
      }

      final String parsedHtmlCode = parseHtmlCode(currUrl);
      final List<String> parsedLinks = parseLinksFromHtmlCode(currUrl, parsedHtmlCode);
      for (String link : parsedLinks) {
        addUrlToUrlsIfIsNotAdded(link);
        if (!processedUrls.contains(link)) {
          tasks.offer(link);
        }
      }

      addUrlToUrlsIfIsNotAdded(currUrl);
      processedUrls.add(currUrl);

      repeatTimes--;
    } while (!tasks.isEmpty()
        && repeatTimes > 0);
  }

  private void addUrlToUrlsIfIsNotAdded(final String url) {
    String urlTitle;
    if (!urls.contains(url)) {
      urlTitle = parseTitle(url);
      urls.add(url);
      urlsTitles.add(urlTitle);
    }
  }
}
