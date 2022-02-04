package maps_test

import (
	"encoding/json"
	"fmt"
	"go-mastery/pkg/maps"
	"testing"
)

func TestNestedLookup(t *testing.T) {
	testCases := []struct {
		desc string
		json string
	}{
		{
			desc: "valid event record",
			json: `{
				"key": {
					"id": "a63f78f9-5679-4bf7-aeab-e2ad0395c998",
					"service": "ASRaaS"
				},
				"offset": 47648049,
				"partition": 0,
				"topic": "Rakuten-VA-Assistant-Prod",
				"value": {
					"appid": "Rakuten-VA-Assistant-Prod",
					"data": {
						"asrSessionid": "a63f78f9-5679-4bf7-aeab-e2ad0395c998",
						"callsummary": {
							"IP!S3Events": "",
							"IPKey": "",
							"Recognize": [
								{
									"absEndTime": 1643585706706,
									"absStartTime": 1643585700766,
									"audioPacketStats": {
										"audioDurationMs": 5840,
										"bytesReceived": 186880,
										"firstPacketTime": "2022-01-30T23:35:00.948Z",
										"lastPacketTime": "2022-01-30T23:35:06.677Z",
										"maximumSize": 640,
										"minimumSize": 640,
										"packetsReceived": 292
									},
									"metrics": {
										"addedWords": {
											"dlm": {
												"1": 87
											}
										},
										"firstAudioToResult": 5757,
										"firstAudioToStartOfSpeech": 392,
										"weights": {
											"dlm": {
												"1": 250
											}
										}
									},
									"recognitionObjects": [
										{
											"S3Id": "DOMAIN_LM1",
											"id": "1",
											"type": "DLM",
											"url": "http://minio-nuance-minio-01.ns-nu-voiceai-cs:9000/nuance-mix-project/training/models/77a2c18f-0b70-4b5f-bcef-fe520d7f4774/builds/7c05a07b-0361-4d50-8304-646d492f3463/content",
											"weight": "MEDIUM"
										}
									],
									"recognitionParameters": {
										"allowZeroBaseLmWeight": false,
										"audioFormat": "audio/l16;rate=16000",
										"autoEnd": true,
										"autoPunctuation": true,
										"confidenceFloatRepresentation": true,
										"enableNBestOnRejection": false,
										"enablePartialResults": true,
										"enableProfanityFilter": false,
										"enableSpeakerProfileUpdate": true,
										"enableUtteranceDetection": true,
										"immutablePartialResults": false,
										"includeTokenization": true,
										"nBestListLength": 5,
										"noInputTimeout": "2000ms",
										"startRecognitionTimers": false,
										"suppressCallRecording": false,
										"suppressInitialCapitalization": false,
										"utteranceEndSilence": "1100ms",
										"wakeupWordFilter": false
									},
									"utterances": [
										{
											"completionCause": "NO-MATCH",
											"info": {
												"clippingDuration": 0,
												"droppedNonSpeechPackets": 0,
												"droppedSpeechPackets": 0,
												"duration": 5810,
												"finalEnergy": -44.1799,
												"initialEnergy": 24.1058,
												"initialSilence": 5700,
												"level": 9729,
												"meanEnergy": 60.1931,
												"snrEstimate": 4,
												"stereo": false
											},
											"nBest": [],
											"relEndTime": 5810,
											"relStartTime": 0
										}
									]
								}
							],
							"absEndTime": 1643585706774,
							"absStartTime": 1643585700714,
							"clientData": {
								"x-nuance-dialog-seqid": "29",
								"x-nuance-dialog-session-id": "002791ab-73af-4402-9475-7d9be7af81d2"
							},
							"dataPack": {
								"id": "ndp40-gen-jpn-JPN-r2",
								"language": "jpn-JPN",
								"topic": "GEN",
								"version": "4.2.0"
							},
							"recognitionObjects": [],
							"serverData": {
								"arch": "x64",
								"hostname": "krypton-ndp-jpn-4-2-56d4d48d65-7snp2",
								"platform": "linux",
								"versions": {
									"Krypton": "4.9.0-202109031315",
									"MRec": "1.42.100.19416",
									"S3": "12.24.100.04528",
									"S3Config": "4.12.0",
									"TextProc": "15.07.100.02136"
								}
							},
							"sessionId": "a63f78f9-5679-4bf7-aeab-e2ad0395c998"
						},
						"dataContentType": "application/x-nuance-asr-callsummary.v2+json",
						"finalResult": {
							"timestamp": 1643585706706
						},
						"processingTime": {
							"durationMs": 6083,
							"startTime": "2022-01-30T23:35:00.693Z"
						},
						"": {
							"recognitionInitMessage": {
								"clientData": {
									"x-nuance-dialog-seqid": "29",
									"x-nuance-dialog-session-id": "002791ab-73af-4402-9475-7d9be7af81d2",
									"x-nuance-event-log-source": "on-premise"
								},
								"parameters": {
									"audioFormat": {
										"pcm": {
											"sampleRateHz": 16000
										}
									},
									"language": "jpn-JPN",
									"maxHypotheses": 5,
									"noInputTimeoutMs": 2000,
									"recognitionFlags": {
										"autoPunctuate": true,
										"includeTokenization": true,
										"maskLoadFailures": true,
										"stallTimers": true
									},
									"resultType": 1,
									"topic": "GEN",
									"utteranceEndSilenceMs": 1100
								},
								"resources": [
									{
										"externalReference": {
											"type": 3,
											"uri": "urn:nuance:mix/jpn-JPN/Lina_Active_Model/mix.asr"
										}
									}
								]
							}
						},
						"response": {
							"status": {
								"code": 200,
								"message": "Success"
							},
							"timestamp": 1643585706776
						},
						"traceid": "264acd74-9403-4bd1-b682-20c31360bdfc"
					},
					"datacontenttype": "application/json",
					"id": "21483868-8491-4591-9e44-1510c3c429e1",
					"partitionKey": "{\"service\":\"ASRaaS\",\"id\":\"a63f78f9-5679-4bf7-aeab-e2ad0395c998\"}",
					"service": "ASRaaS",
					"source": "nuance.asr.v1.Recognizer/Recognize",
					"specversion": "1.0",
					"timestamp": "2022-01-30T23:35:06.777Z",
					"type": "Recognize"
				}
			}`,
		},
		{
			desc: "invalid event record",
			json: `{
				"key": {
					"id": "a63f78f9-5679-4bf7-aeab-e2ad0395c998",
					"service": "ASRaaS"
				},
				"offset": 47648049,
				"partition": 0,
				"topic": "Rakuten-VA-Assistant-Prod",
				"value": {
					"appid": "Rakuten-VA-Assistant-Prod",
					"data": {
						"asrSessionid": "a63f78f9-5679-4bf7-aeab-e2ad0395c998",
						"callsummary": {
							"IP!S3Events": "",
							"IPKey": "",
							"Recognize": [
								{
									"absEndTime": 1643585706706,
									"absStartTime": 1643585700766,
									"audioPacketStats": {
										"audioDurationMs": 5840,
										"bytesReceived": 186880,
										"firstPacketTime": "2022-01-30T23:35:00.948Z",
										"lastPacketTime": "2022-01-30T23:35:06.677Z",
										"maximumSize": 640,
										"minimumSize": 640,
										"packetsReceived": 292
									},
									"metrics": {
										"addedWords": {
											"dlm": {
												"1": 87
											}
										},
										"firstAudioToResult": 5757,
										"firstAudioToStartOfSpeech": 392,
										"weights": {
											"dlm": {
												"1": 250
											}
										}
									},
									"recognitionObjects": [
										{
											"S3Id": "DOMAIN_LM1",
											"id": "1",
											"type": "DLM",
											"url": "http://minio-nuance-minio-01.ns-nu-voiceai-cs:9000/nuance-mix-project/training/models/77a2c18f-0b70-4b5f-bcef-fe520d7f4774/builds/7c05a07b-0361-4d50-8304-646d492f3463/content",
											"weight": "MEDIUM"
										}
									],
									"recognitionParameters": {
										"allowZeroBaseLmWeight": false,
										"audioFormat": "audio/l16;rate=16000",
										"autoEnd": true,
										"autoPunctuation": true,
										"confidenceFloatRepresentation": true,
										"enableNBestOnRejection": false,
										"enablePartialResults": true,
										"enableProfanityFilter": false,
										"enableSpeakerProfileUpdate": true,
										"enableUtteranceDetection": true,
										"immutablePartialResults": false,
										"includeTokenization": true,
										"nBestListLength": 5,
										"noInputTimeout": "2000ms",
										"startRecognitionTimers": false,
										"suppressCallRecording": false,
										"suppressInitialCapitalization": false,
										"utteranceEndSilence": "1100ms",
										"wakeupWordFilter": false
									},
									"utterances": [
										{
											"completionCause": "NO-MATCH",
											"info": {
												"clippingDuration": 0,
												"droppedNonSpeechPackets": 0,
												"droppedSpeechPackets": 0,
												"duration": 5810,
												"finalEnergy": -44.1799,
												"initialEnergy": 24.1058,
												"initialSilence": 5700,
												"level": 9729,
												"meanEnergy": 60.1931,
												"snrEstimate": 4,
												"stereo": false
											},
											"nBest": [],
											"relEndTime": 5810,
											"relStartTime": 0
										}
									]
								}
							],
							"absEndTime": 1643585706774,
							"absStartTime": 1643585700714,
							"": {
								"x-nuance-dialog-seqid": "29",
								"x-nuance-dialog-session-id": "002791ab-73af-4402-9475-7d9be7af81d2"
							},
							"dataPack": {
								"id": "ndp40-gen-jpn-JPN-r2",
								"language": "jpn-JPN",
								"topic": "GEN",
								"version": "4.2.0"
							},
							"recognitionObjects": [],
							"serverData": {
								"arch": "x64",
								"hostname": "krypton-ndp-jpn-4-2-56d4d48d65-7snp2",
								"platform": "linux",
								"versions": {
									"Krypton": "4.9.0-202109031315",
									"MRec": "1.42.100.19416",
									"S3": "12.24.100.04528",
									"S3Config": "4.12.0",
									"TextProc": "15.07.100.02136"
								}
							},
							"sessionId": "a63f78f9-5679-4bf7-aeab-e2ad0395c998"
						},
						"dataContentType": "application/x-nuance-asr-callsummary.v2+json",
						"finalResult": {
							"timestamp": 1643585706706
						},
						"processingTime": {
							"durationMs": 6083,
							"startTime": "2022-01-30T23:35:00.693Z"
						},
						"": {
							"recognitionInitMessage": {
								"": {
									"x-nuance-dialog-seqid": "29",
									"x-nuance-dialog-session-id": "002791ab-73af-4402-9475-7d9be7af81d2",
									"x-nuance-event-log-source": "on-premise"
								},
								"parameters": {
									"audioFormat": {
										"pcm": {
											"sampleRateHz": 16000
										}
									},
									"language": "jpn-JPN",
									"maxHypotheses": 5,
									"noInputTimeoutMs": 2000,
									"recognitionFlags": {
										"autoPunctuate": true,
										"includeTokenization": true,
										"maskLoadFailures": true,
										"stallTimers": true
									},
									"resultType": 1,
									"topic": "GEN",
									"utteranceEndSilenceMs": 1100
								},
								"resources": [
									{
										"externalReference": {
											"type": 3,
											"uri": "urn:nuance:mix/jpn-JPN/Lina_Active_Model/mix.asr"
										}
									}
								]
							}
						},
						"response": {
							"status": {
								"code": 200,
								"message": "Success"
							},
							"timestamp": 1643585706776
						},
						"traceid": "264acd74-9403-4bd1-b682-20c31360bdfc"
					},
					"datacontenttype": "application/json",
					"id": "21483868-8491-4591-9e44-1510c3c429e1",
					"partitionKey": "{\"service\":\"ASRaaS\",\"id\":\"a63f78f9-5679-4bf7-aeab-e2ad0395c998\"}",
					"service": "ASRaaS",
					"source": "nuance.asr.v1.Recognizer/Recognize",
					"specversion": "1.0",
					"timestamp": "2022-01-30T23:35:06.777Z",
					"type": "Recognize"
				}
			}`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var m = make(map[string]interface{})
			json.Unmarshal([]byte(tC.json), &m)
			val, err := maps.NestedMapLookup(m, "value", "data", "callsummary", "clientData")
			if err != nil {
				t.Errorf(err.Error())
			}
			fmt.Printf("(%T) %v\n", val, val)
		})
	}
}
