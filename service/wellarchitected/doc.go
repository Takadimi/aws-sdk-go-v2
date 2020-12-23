// Code generated by smithy-go-codegen DO NOT EDIT.

// Package wellarchitected provides the API client, operations, and parameter types
// for AWS Well-Architected Tool.
//
// AWS Well-Architected Tool This is the AWS Well-Architected Tool API Reference.
// The AWS Well-Architected Tool API provides programmatic access to the AWS
// Well-Architected Tool (http://aws.amazon.com/well-architected-tool) in the AWS
// Management Console (https://console.aws.amazon.com/wellarchitected). Managing
// workloads:
//
// * CreateWorkload: Define a new workload.
//
// * ListWorkloads: List
// workloads.
//
// * GetWorkload: Get the properties of a workload.
//
// * UpdateWorkload:
// Update the properties of a workload.
//
// * DeleteWorkload: Delete a
// workload.
//
// Managing milestones:
//
// * CreateMilestone: Create a milestone.
//
// *
// ListMilestones: List milestones.
//
// * GetMilestone: Get the properties of a
// milestone.
//
// Managing lenses:
//
// * ListLenses: List the available lenses.
//
// *
// AssociateLenses: Add one or more lenses to a workload.
//
// * DisassociateLenses:
// Remove one or more lenses from a workload.
//
// * ListNotifications: List lens
// notifications for a workload.
//
// * GetLensVersionDifference: Get the differences
// between the version of a lens used in a workload and the latest version.
//
// *
// UpgradeLensReview: Upgrade a workload to use the latest version of a
// lens.
//
// Managing reviews:
//
// * ListLensReviews: List reviews associated with a
// workload.
//
// * GetLensReview: Get pillar and risk data associated with a workload
// review.
//
// * GetLensReviewReport: Get the report associated with a workload
// review.
//
// * UpdateLensReview: Update the notes associated with a workload
// review.
//
// * ListAnswers: List the questions and answers in a workload review.
//
// *
// UpdateAnswer: Update the answer to a specific question in a workload review.
//
// *
// ListLensReviewImprovements: List the improvement plan associated with a workload
// review.
//
// Managing workload shares:
//
// * ListWorkloadShares: List the workload
// shares associated with a workload.
//
// * CreateWorkloadShare: Create a workload
// share.
//
// * UpdateWorkloadShare: Update a workload share.
//
// * DeleteWorkloadShare:
// Delete a workload share.
//
// Managing workload share invitations:
//
// *
// ListShareInvitations: List workload share invitations.
//
// * UpdateShareInvitation:
// Update a workload share invitation.
//
// For information about the AWS
// Well-Architected Tool, see the AWS Well-Architected Tool User Guide
// (https://docs.aws.amazon.com/wellarchitected/latest/userguide/intro.html).
package wellarchitected
