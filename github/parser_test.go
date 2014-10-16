package github

import "testing"

func TestParsingGithubData(t *testing.T) {
	body := `{
  "ref": "refs/heads/master",
  "before": "0bce13039c909032a54960cc84586ce63c9b3361",
  "after": "81896ee16894856de90ef8171ab50060094ae763",
  "created": false,
  "deleted": false,
  "forced": false,
  "base_ref": null,
  "compare": "https://github.com/wkirschbaum/ciprojecttest/compare/0bce13039c90...81896ee16894",
  "commits": [
    {
      "id": "81896ee16894856de90ef8171ab50060094ae763",
      "distinct": true,
      "message": "run 3",
      "timestamp": "2014-10-16T15:32:24+02:00",
      "url": "https://github.com/wkirschbaum/ciprojecttest/commit/81896ee16894856de90ef8171ab50060094ae763",
      "author": {
        "name": "Wilhelm Kirschbaum",
        "email": "wkirschbaum@gmail.com",
        "username": "wkirschbaum"
      },
      "committer": {
        "name": "Wilhelm Kirschbaum",
        "email": "wkirschbaum@gmail.com",
        "username": "wkirschbaum"
      },
      "added": [

      ],
      "removed": [

      ],
      "modified": [
        "README.md"
      ]
    }
  ],
  "head_commit": {
    "id": "81896ee16894856de90ef8171ab50060094ae763",
    "distinct": true,
    "message": "run 3",
    "timestamp": "2014-10-16T15:32:24+02:00",
    "url": "https://github.com/wkirschbaum/ciprojecttest/commit/81896ee16894856de90ef8171ab50060094ae763",
    "author": {
      "name": "Wilhelm Kirschbaum",
      "email": "wkirschbaum@gmail.com",
      "username": "wkirschbaum"
    },
    "committer": {
      "name": "Wilhelm Kirschbaum",
      "email": "wkirschbaum@gmail.com",
      "username": "wkirschbaum"
    },
    "added": [

    ],
    "removed": [

    ],
    "modified": [
      "README.md"
    ]
  },
  "repository": {
    "id": 25298379,
    "name": "ciprojecttest",
    "full_name": "wkirschbaum/ciprojecttest",
    "owner": {
      "name": "wkirschbaum",
      "email": "wkirschbaum@users.noreply.github.com"
    },
    "private": false,
    "html_url": "https://github.com/wkirschbaum/ciprojecttest",
    "description": "",
    "fork": false,
    "url": "https://github.com/wkirschbaum/ciprojecttest",
    "forks_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/forks",
    "keys_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/teams",
    "hooks_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/hooks",
    "issue_events_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/issues/events{/number}",
    "events_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/events",
    "assignees_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/assignees{/user}",
    "branches_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/branches{/branch}",
    "tags_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/tags",
    "blobs_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/languages",
    "stargazers_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/stargazers",
    "contributors_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/contributors",
    "subscribers_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/subscribers",
    "subscription_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/subscription",
    "commits_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/issues/comments/{number}",
    "contents_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/contents/{+path}",
    "compare_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/merges",
    "archive_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/downloads",
    "issues_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/issues{/number}",
    "pulls_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/labels{/name}",
    "releases_url": "https://api.github.com/repos/wkirschbaum/ciprojecttest/releases{/id}",
    "created_at": 1413457467,
    "updated_at": "2014-10-16T11:04:27Z",
    "pushed_at": 1413466372,
    "git_url": "git://github.com/wkirschbaum/ciprojecttest.git",
    "ssh_url": "git@github.com:wkirschbaum/ciprojecttest.git",
    "clone_url": "https://github.com/wkirschbaum/ciprojecttest.git",
    "svn_url": "https://github.com/wkirschbaum/ciprojecttest",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "stargazers": 0,
    "master_branch": "master"
  },
  "pusher": {
    "name": "wkirschbaum",
    "email": "wkirschbaum@users.noreply.github.com"
  },
  "sender": {
    "login": "wkirschbaum",
    "id": 1431483,
    "avatar_url": "https://avatars.githubusercontent.com/u/1431483?v=2",
    "gravatar_id": "",
    "url": "https://api.github.com/users/wkirschbaum",
    "html_url": "https://github.com/wkirschbaum",
    "followers_url": "https://api.github.com/users/wkirschbaum/followers",
    "following_url": "https://api.github.com/users/wkirschbaum/following{/other_user}",
    "gists_url": "https://api.github.com/users/wkirschbaum/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/wkirschbaum/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/wkirschbaum/subscriptions",
    "organizations_url": "https://api.github.com/users/wkirschbaum/orgs",
    "repos_url": "https://api.github.com/users/wkirschbaum/repos",
    "events_url": "https://api.github.com/users/wkirschbaum/events{/privacy}",
    "received_events_url": "https://api.github.com/users/wkirschbaum/received_events",
    "type": "User",
    "site_admin": false
  }
}`
	payload, err := ParseCommit([]byte(body))

	if err != nil {
		t.Fatal(err)
	}

	if payload.Ref != "refs/heads/master" {
		t.Fatal("ref not matching")
	}
}
