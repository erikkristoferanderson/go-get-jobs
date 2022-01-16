import time
import os

from util.reddit_code import check_a_subreddit
from util.send_grid_code import send_email


sender_email = os.environ.get('SENDER_EMAIL')
receiver_email = os.environ.get('RECEIVER_EMAIL')

while True:
    try:
        results = check_a_subreddit('hiring')

        if not len(results) > 0:
            print('nothing to see here')
        else:
            for result in results:
                print(result)
                poster_name = result.split('\n')[1]
                time_stamp = result.split('\n')[2]
                send_email(sender=sender_email,
                           recipient=receiver_email,
                           subject=f'new result from {poster_name} at {time_stamp}',
                           html_content=result.replace('\n', '<br>')
                           )

    except Exception as e:
        print('some kind of error')
        print(e)
    finally:
        time.sleep(600)
