package main

import (
	"reflect"
	"testing"
)

func Test_unmarshal2map(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "test",
			args: args{
				in: "{\"dc\":{\"params\":{},\"result\":{},\"scenes\":[{\"scene_id\":\"3675976609414859\",\"decision_rule_id\":\"\",\"decision_rule_priority\":0,\"decisions\":\"\",\"decision_config\":\"\",\"rules\":[{\"rule_id\":\"349737\",\"hit\":false,\"time_cost\":0,\"start_time\":0,\"end_time\":0,\"priority\":0,\"error\":\"\",\"decisions\":\"MISS\",\"decision_config\":\"\",\"detailed_rule_grey_strategy\":{\"grey_version_dimension\":0,\"is_hit_task\":false,\"no_need_anti_brush\":false,\"no_need_list_task_import\":false,\"task_id\":1000,\"task_type\":1,\"task_execute_type\":1},\"identification_tags\":null,\"execution_module\":0,\"punishes\":null,\"miss_type\":2},{\"rule_id\":\"349737\",\"hit\":false,\"time_cost\":1364,\"start_time\":0,\"end_time\":0,\"priority\":0,\"error\":\"\",\"decisions\":\"MISS\",\"decision_config\":\"\",\"detailed_rule_grey_strategy\":{\"grey_version_dimension\":2,\"is_hit_task\":true,\"no_need_anti_brush\":true,\"no_need_list_task_import\":true,\"task_id\":1000,\"task_type\":1,\"task_execute_type\":1},\"identification_tags\":null,\"execution_module\":0,\"punishes\":null,\"miss_type\":0}],\"execution_modules\":null,\"layer\":0,\"scene_condition_hit\":true}],\"event\":\"wanghuanlin_test\",\"decision_scene_id\":\"\",\"decision_scene_priority\":0,\"sync_decision_scene_id\":\"\",\"sync_decision_scene_priority\":0,\"async_decision_scene_id\":\"\",\"async_decision_scene_priority\":0,\"server_time\":1716707659,\"event_data\":\"\",\"sentry_context\":null,\"event_operation\":\"\",\"version\":0,\"access_detailed_content\":null,\"sync_access_detailed_content\":null,\"double_run_content\":{\"scenes\":null}},\"params\":{\"\":\"\",\"IDC\":\"boe\",\"__caller\":\"\",\"__dataQualityBizReqTime\":1716707659,\"__dataQualityMsgId\":\"20240526151419E5224ADE895BEB42B211:grIeYuKYHQDSYXVfHwWV\",\"__dataQualityMsgTime\":1716707659,\"__decision_scene_exec_config\":{\"status\":1,\"rate\":10000,\"canary_mode\":0,\"punish_mode\":0},\"__factor_exec_time\":{\"f_1028\":{\"self\":26,\"total\":33,\"start_time\":1716707659624315,\"level\":\"0\"},\"f_78861\":{\"self\":748,\"total\":762,\"start_time\":1716707659620987,\"level\":\"0\"}},\"__grey_task_hit_version\":{\"1000\":1},\"__grey_version_no_need_send_audit_rule\":{\"349737\":true},\"__layered_async_dec_diff\":0,\"__layered_async_second_dec_diff\":0,\"__layered_decision_detail\":{\"exec_status\":1,\"identify_layered_final_result\":{\"layer\":0,\"inner_final_result\":{\"scene_id\":\"\",\"rule_id\":\"\",\"decision\":\"MISS\",\"second_decision\":\"\",\"decision_priority\":0},\"final_decision_scene_id\":\"\",\"final_decision_rule_id\":\"\",\"final_punish_scene_id\":\"\",\"final_punish_rule_id\":\"\",\"final_decision\":\"MISS\",\"final_second_decision\":\"\",\"punish_configs\":null,\"punishes\":null},\"decision_layered_final_result\":{\"layer\":1,\"inner_final_result\":{\"scene_id\":\"\",\"rule_id\":\"\",\"decision\":\"MISS\",\"second_decision\":\"\",\"decision_priority\":0},\"final_decision_scene_id\":\"\",\"final_decision_rule_id\":\"\",\"final_punish_scene_id\":\"\",\"final_punish_rule_id\":\"\",\"final_decision\":\"MISS\",\"final_second_decision\":\"\",\"punish_configs\":null,\"punishes\":null},\"tmp_details\":{\"decision_scene_hit_detail\":{},\"scene_hit_tag_mapping\":{\"114379\":{}}}},\"__layered_exec_status\":1,\"__layered_inner_dec_diff\":0,\"__layered_inner_second_dec_diff\":0,\"__layered_punish_diff\":0,\"__process_timeline\":{\"anti_call_inner\":1716707659581,\"anti_start_check\":1716707659509,\"inner_start_check\":1716707659612},\"__ruleplatform_not_write_in_execute\":true,\"__shark_final_decision_info\":{\"final_decision\":\"MISS\",\"final_decision_config\":\"\",\"final_punish_configs\":null},\"aid\":-6,\"comment_id\":\"wanghuanlin\",\"did\":-6,\"docking_type__\":\"RPC\",\"event\":\"wanghuanlin_test\",\"eventTime\":0,\"event_time\":0,\"f_1028\":1238761233,\"f_78861\":\"wanghuanlin\",\"iid\":-6,\"is_login\":false,\"log_id\":\"20240526151419E5224ADE895BEB42B211\",\"namespace_id\":13,\"shark_request_id\":\"021716707659514fdbddc0100fff003ffffffff000001c88cb4a4\",\"uid\":1238761233,\"ut\":12}}",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unmarshal2map(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshal2map() = %v, want %v", got, tt.want)
			}
		})
	}
}
