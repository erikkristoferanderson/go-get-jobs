"""Uses praw to look for posts on reddit
relies on environment variables:
- REDDIT_API_CLIENT_ID
- REDDIT_API_CLIENT_SECRET
- REDDIT_API_USER_AGENT
"""

import os
import praw
from datetime import datetime, timedelta


def check_a_subreddit(subreddit_name):
    results_list = []
    sub_name = subreddit_name
    results_limit = 20

    reddit = get_reddit_client()

    for submission in reddit.subreddit(sub_name).new(limit=results_limit):
        if meets_conditions(submission):
            results_list.append(get_text_from_submission(submission))

    return results_list


def meets_conditions(submission, search_terms):
    time_cutoff = datetime.now() - timedelta(minutes=10, seconds=15)
    submission_title = submission.title.lower()
    submission_time = datetime.fromtimestamp(submission.created_utc)

    title_condition = all([submission_title.find(search_term) >= 0 for search_term in search_terms])
    time_condition = submission_time > time_cutoff

    return title_condition and time_condition


def get_text_from_submission(submission):
    result_text = ''
    result_text += submission.subreddit.display_name + ' \n '
    result_text += submission.author.name + ' \n '
    result_text += str(datetime.fromtimestamp(submission.created_utc) - timedelta(hours=5)) + ' \n '
    result_text += submission.title + ' \n '
    result_text += submission.url + ' \n '
    result_text += ' \n '

    return result_text


def get_reddit_client():
    my_client_id = os.environ.get("REDDIT_API_CLIENT_ID")
    my_client_secret = os.environ.get("REDDIT_API_CLIENT_SECRET")
    my_user_agent = os.environ.get("REDDIT_API_USER_AGENT")
    reddit = praw.Reddit(client_id=my_client_id, client_secret=my_client_secret, user_agent=my_user_agent)
    return reddit