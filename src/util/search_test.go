package util_test

import (
	"sort"
	"testing"
	"util"
)

type Campaign struct {
	Id   int
	Name string
}

type ByCampaignId []*Campaign

func (t ByCampaignId) Len() int           { return len(t) }
func (t ByCampaignId) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByCampaignId) Less(i, j int) bool { return t[i].Id < t[j].Id }

func (t ByCampaignId) Compare(i int, element interface{}) int {
	campaignId := element.(int)

	if campaignId == t[i].Id {
		return 0
	}

	if campaignId > t[i].Id {
		return 1
	} else {
		return -1
	}
}

func TestBinarySearch(t *testing.T) {
	// create n campaign
	campaigns := make([]*Campaign, 0)
	campaign1 := &Campaign{Id: 2, Name: "C2"}
	campaign2 := &Campaign{Id: 7, Name: "C7"}
	campaign3 := &Campaign{Id: 5, Name: "C5"}
	campaign4 := &Campaign{Id: 3, Name: "C3"}
	campaign5 := &Campaign{Id: 4, Name: "C4"}
	campaign6 := &Campaign{Id: 6, Name: "C6"}
	campaign7 := &Campaign{Id: 1, Name: "C1"}

	campaigns = append(campaigns, campaign1)
	campaigns = append(campaigns, campaign2)
	campaigns = append(campaigns, campaign3)
	campaigns = append(campaigns, campaign4)
	campaigns = append(campaigns, campaign5)
	campaigns = append(campaigns, campaign6)
	campaigns = append(campaigns, campaign7)

	/*
		t.Log("Before sort")
		for idx, campaign := range campaigns {
			t.Logf("idx: %d %v\n", idx, *campaign)
		}
	*/
	// sort campain first in asending order
	sort.Sort(ByCampaignId(campaigns))
	/*
		t.Log("After sort")
		for idx, campaign := range campaigns {
			t.Logf("idx: %d %v\n", idx, *campaign)
		}
	*/

	founIdx := util.BinarySearch(ByCampaignId(campaigns), 1)
	if 0 != founIdx {
		t.Fatal("Invalid found index is returned")
	}

	founIdx = util.BinarySearch(ByCampaignId(campaigns), 7)
	if 6 != founIdx {
		t.Fatal("Invalid found index is returned")
	}

	founIdx = util.BinarySearch(ByCampaignId(campaigns), 4)
	if 3 != founIdx {
		t.Fatal("Invalid found index is returned")
	}

	founIdx = util.BinarySearch(ByCampaignId(campaigns), 3)
	if 2 != founIdx {
		t.Fatal("Invalid found index is returned")
	}

	founIdx = util.BinarySearch(ByCampaignId(campaigns), 10)
	if -1 != founIdx {
		t.Fatal("Invalid found index is returned")
	}
}
