
import json
import datetime
import re
import argparse

def create_slug(title):
    slug = title.lower()
    slug = slug.replace(' ', '-')
    slug = re.sub(r'[^a-zA-Z0-9-]', '', slug)
    return slug

def add_post(args):
    """Creates a new blog post from command-line arguments and adds it to allposts.json."""

    # --- Read Markdown Content ---
    try:
        with open(args.content_path, 'r', encoding='utf-8') as f:
            content = f.read()
    except FileNotFoundError:
        print(f"Error: The file '{args.content_path}' was not found.")
        return

    # --- Load Existing Posts ---
    try:
        with open('allposts.json', 'r', encoding='utf-8') as f:
            posts = json.load(f)
    except FileNotFoundError:
        posts = []

    # --- Generate New Post Object ---
    new_id = str(int(posts[-1]['id']) + 1) if posts else "1"
    now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    new_post = {
        "id": new_id,
        "slug": create_slug(args.title),
        "title": args.title,
        "subtitle": args.subtitle,
        "author": args.author,
        "content": content,
        "summary": args.summary,
        "read_time": args.read_time,
        "tags": args.tags,
        "category": args.category,
        "created_on": now,
        "updated_on": now,
        "published": "1" if args.published else "0",
        "featured": "1" if args.featured else "0"
    }

    # --- Add New Post and Save ---
    posts.append(new_post)

    with open('allposts.json', 'w', encoding='utf-8') as f:
        json.dump(posts, f, indent=2, ensure_ascii=False)

    print(f"\nSuccessfully added new post with title: '{args.title}'")
    print(f"The new post has been saved to allposts.json.")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Add a new blog post to allposts.json.")
    parser.add_argument("--title", required=True, help="The title of the post.")
    parser.add_argument("--subtitle", required=True, help="The subtitle of the post.")
    parser.add_argument("--author", required=True, help="The author of the post.")
    parser.add_argument("--summary", required=True, help="A short summary of the post.")
    parser.add_argument("--read-time", required=True, help="The estimated read time.")
    parser.add_argument("--tags", required=True, help="Comma-separated tags for the post.")
    parser.add_argument("--category", required=True, help="The category of the post.")
    parser.add_argument("--content-path", required=True, help="The path to the markdown file for the content.")
    parser.add_argument("--published", action='store_true', help="Set if the post is published.")
    parser.add_argument("--featured", action='store_true', help="Set if the post is featured.")

    args = parser.parse_args()
    add_post(args)
