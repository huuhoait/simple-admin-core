import { defHttp } from '/@/utils/http/axios';
import { ErrorMessageMode } from '/#/axios';
import { BaseDataResp, BaseListReq, BaseResp, BaseIDsReq, BaseIDReq } from '/@/api/model/baseModel';
import { AuditInfo, AuditListResp } from './model/auditModel';

enum Api {
  CreateAudit = '/sys-api/audit/create',
  UpdateAudit = '/sys-api/audit/update',
  GetAuditList = '/sys-api/audit/list',
  DeleteAudit = '/sys-api/audit/delete',
  GetAuditById = '/sys-api/audit',
}

/**
 * @description: Get audit list
 */

export const getAuditList = (params: BaseListReq, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseDataResp<AuditListResp>>(
    { url: Api.GetAuditList, params },
    { errorMessageMode: mode },
  );
};

/**
 *  @description: Create a new audit
 */
export const createAudit = (params: AuditInfo, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseResp>(
    { url: Api.CreateAudit, params: params },
    {
      errorMessageMode: mode,
    },
  );
};

/**
 *  @description: Update the audit
 */
export const updateAudit = (params: AuditInfo, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseResp>(
    { url: Api.UpdateAudit, params: params },
    {
      errorMessageMode: mode,
    },
  );
};

/**
 *  @description: Delete audits
 */
export const deleteAudit = (params: BaseIDsReq, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseResp>(
    { url: Api.DeleteAudit, params: params },
    {
      errorMessageMode: mode,
    },
  );
};

/**
 *  @description: Get audit By ID
 */
export const getAuditById = (params: BaseIDReq, mode: ErrorMessageMode = 'message') => {
  return defHttp.post<BaseDataResp<AuditInfo>>(
    { url: Api.GetAuditById, params: params },
    {
      errorMessageMode: mode,
    },
  );
};
