import { defHttp } from '/@/utils/http/axios';
import { ErrorMessageMode } from '/#/axios';
import { BaseDataResp, BaseListReq, BaseResp, BaseIDsReq, BaseIDReq } from '/@/api/model/baseModel';
import { MemberRankInfo, MemberRankListResp } from './model/memberRankModel';

enum Api {
  CreateMemberRank = '/example-api/member_rank/create',
  UpdateMemberRank = '/example-api/member_rank/update',
  GetMemberRankList = '/example-api/member_rank/list',
  DeleteMemberRank = '/example-api/member_rank/delete',
  GetMemberRankById = '/example-api/member_rank',
}

/**
 * @description: Get member rank list
 */

export const getMemberRankList = (params: BaseListReq, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseDataResp<MemberRankListResp>>(
    { url: Api.GetMemberRankList, params },
    { errorMessageMode: mode },
  );
};

/**
 *  @description: Create a new member rank
 */
export const createMemberRank = (params: MemberRankInfo, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseResp>(
    { url: Api.CreateMemberRank, params: params },
    {
      errorMessageMode: mode,
    },
  );
};

/**
 *  @description: Update the member rank
 */
export const updateMemberRank = (params: MemberRankInfo, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseResp>(
    { url: Api.UpdateMemberRank, params: params },
    {
      errorMessageMode: mode,
    },
  );
};

/**
 *  @description: Delete member ranks
 */
export const deleteMemberRank = (params: BaseIDsReq, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseResp>(
    { url: Api.DeleteMemberRank, params: params },
    {
      errorMessageMode: mode,
    },
  );
};

/**
 *  @description: Get member rank By ID
 */
export const getMemberRankById = (params: BaseIDReq, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseDataResp<MemberRankInfo>>(
    { url: Api.GetMemberRankById, params: params },
    {
      errorMessageMode: mode,
    },
  );
};
