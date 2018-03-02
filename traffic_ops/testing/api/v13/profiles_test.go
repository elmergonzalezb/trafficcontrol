/*

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package v13

import (
	"fmt"
	"testing"

	"github.com/apache/incubator-trafficcontrol/lib/go-log"
	tc "github.com/apache/incubator-trafficcontrol/lib/go-tc"
)

func TestProfiles(t *testing.T) {

	CreateTestCDNs(t)
	CreateTestProfiles(t)
	UpdateTestProfiles(t)
	GetTestProfiles(t)
	DeleteTestProfiles(t)
	DeleteTestCDNs(t)

}

func CreateTestProfiles(t *testing.T) {

	for _, pr := range testData.Profiles {
		cdns, _, err := TOSession.GetCDNByName(pr.CDNName)
		fmt.Printf("cdns ---> %v\n", cdns)
		respCDN := cdns[0]
		pr.CDNName = respCDN.Name

		resp, _, err := TOSession.CreateProfile(pr)
		log.Debugln("Response: ", resp)
		if err != nil {
			t.Errorf("could not CREATE phys_locations: %v\n", err)
		}
	}

}

func UpdateTestProfiles(t *testing.T) {

	firstProfile := testData.Profiles[0]
	// Retrieve the Profile by name so we can get the id for the Update
	resp, _, err := TOSession.GetProfileByName(firstProfile.Name)
	if err != nil {
		t.Errorf("cannot GET Profile by name: %v - %v\n", firstProfile.Name, err)
	}
	remoteProfile := resp[0]
	expectedProfileName := "UPDATED"
	remoteProfile.Name = expectedProfileName
	var alert tc.Alerts
	alert, _, err = TOSession.UpdateProfileByID(remoteProfile.ID, remoteProfile)
	if err != nil {
		t.Errorf("cannot UPDATE Profile by id: %v - %v\n", err, alert)
	}

	// Retrieve the Profile to check Profile name got updated
	resp, _, err = TOSession.GetProfileByID(remoteProfile.ID)
	if err != nil {
		t.Errorf("cannot GET Profile by name: %v - %v\n", firstProfile.Name, err)
	}
	respProfile := resp[0]
	if respProfile.Name != expectedProfileName {
		t.Errorf("results do not match actual: %s, expected: %s\n", respProfile.Name, expectedProfileName)
	}

}

func GetTestProfiles(t *testing.T) {

	for _, pr := range testData.Profiles {
		resp, _, err := TOSession.GetProfileByName(pr.Name)
		if err != nil {
			t.Errorf("cannot GET Profile by name: %v - %v\n", err, resp)
		}
	}
}

func DeleteTestProfiles(t *testing.T) {

	for _, pr := range testData.Profiles {
		// Retrieve the Profile by name so we can get the id for the Update
		resp, _, err := TOSession.GetProfileByName(pr.Name)
		if err != nil {
			t.Errorf("cannot GET Profile by name: %v - %v\n", pr.Name, err)
		}
		if len(resp) > 0 {
			respProfile := resp[0]

			delResp, _, err := TOSession.DeleteProfileByID(respProfile.ID)
			if err != nil {
				t.Errorf("cannot DELETE Profile by name: %v - %v\n", err, delResp)
			}
			//time.Sleep(1 * time.Second)

			// Retrieve the Profile to see if it got deleted
			prs, _, err := TOSession.GetProfileByName(pr.Name)
			if err != nil {
				t.Errorf("error deleting Profile name: %s\n", err.Error())
			}
			if len(prs) > 0 {
				t.Errorf("expected Profile Name: %s to be deleted\n", pr.Name)
			}
		} else {
			fmt.Printf("no resp ---> %v\n", resp)
		}
	}
}