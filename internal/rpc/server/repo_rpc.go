package rpc_server

import (
	. "gogs.io/gogs/internal/db"
	shared "gogs.io/gogs/internal/rpc/shared"
	log "unknwon.dev/clog/v2"
)

type Repo int

func (t *Repo) DeleteRepository(args *shared.Repo_DeleteRepositoryArgs, reply *interface{}) error {
	log.Info("Legacy %v", HasEngine)
	return nil

	// // Extract the values passed in `args`
	// repoID := args.RepoID
	// ownerID := args.OwnerID

	// repo := &Repository{ID: repoID, OwnerID: ownerID}
	// has, err := x.Get(repo)
	// if err != nil {
	// 	return err
	// } else if !has {
	// 	return ErrRepoNotExist{args: map[string]interface{}{"ownerID": ownerID, "repoID": repoID}}
	// }

	// // In case is a organization.
	// org, err := GetUserByID(ownerID)
	// if err != nil {
	// 	return err
	// }
	// if org.IsOrganization() {
	// 	if err = org.GetTeams(); err != nil {
	// 		return err
	// 	}
	// }

	// sess := x.NewSession()
	// defer sess.Close()
	// if err = sess.Begin(); err != nil {
	// 	return err
	// }

	// if org.IsOrganization() {
	// 	for _, t := range org.Teams {
	// 		if !t.hasRepository(sess, repoID) {
	// 			continue
	// 		} else if err = t.removeRepository(sess, repo, false); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// if err = deleteBeans(sess,
	// 	&Repository{ID: repoID},
	// 	&Access{RepoID: repo.ID},
	// 	&Action{RepoID: repo.ID},
	// 	&Watch{RepoID: repoID},
	// 	&Star{RepoID: repoID},
	// 	&Mirror{RepoID: repoID},
	// 	&IssueUser{RepoID: repoID},
	// 	&Milestone{RepoID: repoID},
	// 	&Release{RepoID: repoID},
	// 	&Collaboration{RepoID: repoID},
	// 	&PullRequest{BaseRepoID: repoID},
	// 	&ProtectBranch{RepoID: repoID},
	// 	&ProtectBranchWhitelist{RepoID: repoID},
	// 	&Webhook{RepoID: repoID},
	// 	&HookTask{RepoID: repoID},
	// 	&LFSObject{RepoID: repoID},
	// ); err != nil {
	// 	return fmt.Errorf("deleteBeans: %v", err)
	// }

	// // Delete comments and attachments.
	// issues := make([]*Issue, 0, 25)
	// attachmentPaths := make([]string, 0, len(issues))
	// if err = sess.Where("repo_id=?", repoID).Find(&issues); err != nil {
	// 	return err
	// }
	// for i := range issues {
	// 	if _, err = sess.Delete(&Comment{IssueID: issues[i].ID}); err != nil {
	// 		return err
	// 	}

	// 	attachments := make([]*Attachment, 0, 5)
	// 	if err = sess.Where("issue_id=?", issues[i].ID).Find(&attachments); err != nil {
	// 		return err
	// 	}
	// 	for j := range attachments {
	// 		attachmentPaths = append(attachmentPaths, attachments[j].LocalPath())
	// 	}

	// 	if _, err = sess.Delete(&Attachment{IssueID: issues[i].ID}); err != nil {
	// 		return err
	// 	}
	// }

	// if _, err = sess.Delete(&Issue{RepoID: repoID}); err != nil {
	// 	return err
	// }

	// if repo.IsFork {
	// 	if _, err = sess.Exec("UPDATE `repository` SET num_forks=num_forks-1 WHERE id=?", repo.ForkID); err != nil {
	// 		return fmt.Errorf("decrease fork count: %v", err)
	// 	}
	// }

	// if _, err = sess.Exec("UPDATE `user` SET num_repos=num_repos-1 WHERE id=?", ownerID); err != nil {
	// 	return err
	// }

	// if err = sess.Commit(); err != nil {
	// 	return fmt.Errorf("Commit: %v", err)
	// }

	// // Remove repository files.
	// repoPath := repo.RepoPath()
	// RemoveAllWithNotice("Delete repository files", repoPath)

	// repo.DeleteWiki()

	// // Remove attachment files.
	// for i := range attachmentPaths {
	// 	RemoveAllWithNotice("Delete attachment", attachmentPaths[i])
	// }

	// if repo.NumForks > 0 {
	// 	if _, err = x.Exec("UPDATE `repository` SET fork_id=0,is_fork=? WHERE fork_id=?", false, repo.ID); err != nil {
	// 		log.Error("reset 'fork_id' and 'is_fork': %v", err)
	// 	}
	// }

	// return nil
}
