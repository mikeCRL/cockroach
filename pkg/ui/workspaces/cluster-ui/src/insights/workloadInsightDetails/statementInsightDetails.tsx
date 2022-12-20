// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.
import React, { useEffect, useState } from "react";
import Helmet from "react-helmet";
import { RouteComponentProps } from "react-router-dom";
import { ArrowLeft } from "@cockroachlabs/icons";
import { Row, Col, Tabs } from "antd";
import "antd/lib/tabs/style";
import "antd/lib/col/style";
import "antd/lib/row/style";
import { Button } from "src/button";
import { Loading } from "src/loading";
import { SqlBox, SqlBoxSize } from "src/sql";
import { getMatchParamByName, idAttr } from "src/util";
import { StmtInsightEvent } from "../types";
import { InsightsError } from "../insightsErrorComponent";
import { getExplainPlanFromGist } from "src/api/decodePlanGistApi";
import { StatementInsightDetailsOverviewTab } from "./statementInsightDetailsOverviewTab";
import { executionInsightsRequestFromTimeScale } from "../utils";
import { TimeScale } from "../../timeScaleDropdown";
import { StmtInsightsReq } from "src/api";
import LoadingError from "../../sqlActivity/errorComponent";

// Styles
import classNames from "classnames/bind";
import { commonStyles } from "src/common";
import insightsDetailsStyles from "src/insights/workloadInsightDetails/insightsDetails.module.scss";

const cx = classNames.bind(insightsDetailsStyles);

enum TabKeysEnum {
  OVERVIEW = "overview",
  EXPLAIN = "explain",
}
export interface StatementInsightDetailsStateProps {
  insightEventDetails: StmtInsightEvent;
  insightError: Error | null;
  isTenant?: boolean;
  timeScale?: TimeScale;
  hasAdminRole: boolean;
}

export interface StatementInsightDetailsDispatchProps {
  refreshStatementInsights: (req: StmtInsightsReq) => void;
  setTimeScale: (ts: TimeScale) => void;
  refreshUserSQLRoles: () => void;
}

export type StatementInsightDetailsProps = StatementInsightDetailsStateProps &
  StatementInsightDetailsDispatchProps &
  RouteComponentProps<unknown>;

type ExplainPlanState = {
  explainPlan: string;
  loaded: boolean;
  error: Error;
};

export const StatementInsightDetails: React.FC<
  StatementInsightDetailsProps
> = ({
  history,
  insightEventDetails,
  insightError,
  match,
  isTenant,
  timeScale,
  hasAdminRole,
  setTimeScale,
  refreshStatementInsights,
  refreshUserSQLRoles,
}) => {
  const [explainPlanState, setExplainPlanState] = useState<ExplainPlanState>({
    explainPlan: null,
    loaded: false,
    error: null,
  });

  const prevPage = (): void => history.goBack();

  const onTabClick = (key: TabKeysEnum) => {
    if (
      !isTenant &&
      key === TabKeysEnum.EXPLAIN &&
      insightEventDetails?.planGist &&
      !explainPlanState.loaded
    ) {
      // Get the explain plan.
      getExplainPlanFromGist({ planGist: insightEventDetails.planGist }).then(
        res => {
          setExplainPlanState({
            explainPlan: res.explainPlan,
            loaded: true,
            error: res.error,
          });
        },
      );
    }
  };

  const executionID = getMatchParamByName(match, idAttr);

  useEffect(() => {
    refreshUserSQLRoles();
    if (!insightEventDetails || insightEventDetails === null) {
      const req = executionInsightsRequestFromTimeScale(timeScale);
      refreshStatementInsights(req);
    }
  }, [
    insightEventDetails,
    timeScale,
    refreshStatementInsights,
    refreshUserSQLRoles,
  ]);

  return (
    <div>
      <Helmet title={"Details | Insight"} />
      <Button
        onClick={prevPage}
        type="unstyled-link"
        size="small"
        icon={<ArrowLeft fontSize={"10px"} />}
        iconPosition="left"
        className={commonStyles("small-margin")}
      >
        Insights
      </Button>
      <h3 className={commonStyles("base-heading", "no-margin-bottom")}>
        Statement Execution ID: {executionID}
      </h3>
      <div>
        <Loading
          loading={insightEventDetails === null}
          page={"Statement Insight details"}
          error={insightError}
          renderError={() => InsightsError()}
        >
          <section className={cx("section")}>
            <Row>
              <Col span={24}>
                <SqlBox
                  size={SqlBoxSize.custom}
                  value={insightEventDetails?.query}
                />
              </Col>
            </Row>
          </section>
          <Tabs
            className={commonStyles("cockroach--tabs")}
            defaultActiveKey={TabKeysEnum.OVERVIEW}
            onTabClick={onTabClick}
          >
            <Tabs.TabPane tab="Overview" key={TabKeysEnum.OVERVIEW}>
              <StatementInsightDetailsOverviewTab
                insightEventDetails={insightEventDetails}
                setTimeScale={setTimeScale}
                hasAdminRole={hasAdminRole}
              />
            </Tabs.TabPane>
            {!isTenant && (
              <Tabs.TabPane tab="Explain Plan" key={TabKeysEnum.EXPLAIN}>
                <section className={cx("section")}>
                  <Row gutter={24}>
                    <Col span={24}>
                      <Loading
                        loading={
                          !explainPlanState.loaded &&
                          insightEventDetails?.planGist?.length > 0
                        }
                        page={"stmt_insight_details"}
                        error={explainPlanState.error}
                        renderError={() =>
                          LoadingError({
                            statsType: "explain plan",
                            timeout: explainPlanState.error?.name
                              ?.toLowerCase()
                              .includes("timeout"),
                          })
                        }
                      >
                        <SqlBox
                          value={
                            explainPlanState.explainPlan || "Not available."
                          }
                          size={SqlBoxSize.custom}
                        />
                      </Loading>
                    </Col>
                  </Row>
                </section>
              </Tabs.TabPane>
            )}
          </Tabs>
        </Loading>
      </div>
    </div>
  );
};
