package repository

import (
	"context"
	"log"
	"time"
	"math/rand"
	"strconv"

	t "github.com/PyMarcus/codes_download/tools"
)

func Insert(jsonPath string) {
    var itemId int64

    conn := psqlConnect()
    log.Println("reading json file and setting up structs")
    topics, owners, licenses, items, imp := t.ReadJsonFile(jsonPath)

    // Start the transaction
    tx, err := conn.Begin(context.Background())
    if err != nil {
        log.Fatal("failed to begin transaction:", err)
    }

    defer func() {
        if err != nil {
            log.Println("fail to insert data", err)
            // log.Println("rolling back transaction due to error:", err)
            // tx.Rollback(context.Background())
        } else {
            log.Println("OK! Trying to commit")
            err = tx.Commit(context.Background())
            if err != nil {
                log.Println("failed to commit transaction:", err)
            }
        }
        conn.Close()
    }()
    
	var idImport int64
	if err := tx.QueryRow(
		context.Background(),
		"INSERT INTO tb_imports (total_count, incomplete_results) VALUES ($1, $2) RETURNING id_import",
		imp.TotalCount, false,
	).Scan(&idImport); err != nil {
		log.Println("error inserting imports:", err)
		return
	}

    log.Println("inserting items")
    for _, item := range items {
		query := `
        INSERT INTO tb_items 
        VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, 
            $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, 
            $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, 
            $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, 
            $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, 
            $73, $74, $75, $76, $77, $78, $79, $80, $81
        ) RETURNING id_item`

    if err := tx.QueryRow(
        context.Background(), query,
        item.Id, idImport, strconv.Itoa(item.Id), item.NodeID, item.Name, item.FullName,
        item.Private, item.OwnerLogin, item.OwnerID, item.HTMLURL, item.Description,
        item.Fork, item.URL, item.ForksURL, item.KeysURL, item.CollaboratorsURL,
        item.TeamsURL, item.HooksURL, item.IssueEventsURL, item.EventsURL,
        item.AssigneesURL, item.BranchesURL, item.TagsURL, item.BlobsURL,
        item.GitTagsURL, item.GitRefsURL, item.TreesURL, item.StatusesURL,
        item.LanguagesURL, item.StargazersURL, item.ContributorsURL,
        item.SubscribersURL, item.SubscriptionURL, item.CommitsURL,
        item.GitCommitsURL, item.CommentsURL, item.IssueCommentURL,
        item.ContentsURL, item.CompareURL, item.MergesURL, item.ArchiveURL,
        item.DownloadsURL, item.IssuesURL, item.PullsURL, item.MilestonesURL,
        item.NotificationsURL, item.LabelsURL, item.ReleasesURL,
        item.DeploymentsURL, item.CreatedAt, item.UpdatedAt, item.PushedAt,
        item.GitURL, item.SSHURL, item.CloneURL, item.SVNURL, item.Homepage,
        strconv.Itoa(item.Size), strconv.Itoa(item.StargazersCount), strconv.Itoa(item.WatchersCount), item.Language,
        item.HasIssues, item.HasProjects, item.HasDownloads, item.HasWiki,
        item.HasPages, item.HasDiscussions, strconv.Itoa(item.ForksCount), item.MirrorURL,
        item.Archived, item.Disabled, strconv.Itoa(item.OpenIssuesCount), item.AllowForking,
        item.IsTemplate, item.WebCommitSignoffReqd, item.Visibility, strconv.Itoa(item.Forks),
        strconv.Itoa(item.OpenIssues), strconv.Itoa(item.Watchers), item.DefaultBranch, strconv.Itoa(item.Score),
    ).Scan(&itemId); err != nil {
        log.Println("error inserting items:", err)
        return
    }

    }
    
    
    rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000000)

    log.Println("inserting topics")
    for _, topic := range topics {
        query := `
		INSERT INTO tb_topics (
			id_topic, id_item, topic_name
		) 
		VALUES (
			$1, $2, $3
		)`
		randomNumber = rand.Intn(1000000000)

        if _, err := tx.Exec(context.Background(), query,
		randomNumber, itemId, topic.TopicName); err != nil {
            log.Println("error inserting topics:", err)
            return
        }
    }

    log.Println("inserting owners")
    for _, owner := range owners {
        query := `
		INSERT INTO tb_owner (
			id_owner, id_item, login, id, node_id, avatar_url, gravatar_id, 
			url, html_url, followers_url, following_url, gists_url, starred_url, 
			subscriptions_url, organizations_url, repos_url, events_url, 
			received_events_url, type, site_admin
		) 
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, 
			$17, $18, $19, $20
		)`
		randomNumber = rand.Intn(1000000000)

        if _, err := tx.Exec(context.Background(), query,
		randomNumber, itemId, owner.Login, strconv.Itoa(owner.Id), owner.NodeID,
            owner.AvatarURL, owner.GravatarID, owner.URL, owner.HTMLURL,
            owner.FollowersURL, owner.FollowingURL, owner.GistsURL, owner.StarredURL,
            owner.SubscriptionsURL, owner.OrganizationsURL, owner.ReposURL,
            owner.EventsURL, owner.ReceivedEventsURL, owner.Type, strconv.FormatBool(owner.SiteAdmin)); err != nil {
            log.Println("error inserting owners:", err)
            return
        }
    }

    log.Println("inserting licenses")
    for _, license := range licenses {
        query := `
		INSERT INTO tb_license (
			id_license, id_item, key, name, spdx_id, url, node_id
		) 
		VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)`
		randomNumber = rand.Intn(1000000000)

        if _, err := tx.Exec(context.Background(), query,
		randomNumber, itemId, license.Key, license.Name,
            license.SPDXID, license.URL, license.NodeID); err != nil {
            log.Println("error inserting licenses:", err)
            return
        }
    }
}
