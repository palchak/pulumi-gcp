// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class TableExternalDataConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - Let BigQuery try to autodetect the schema
        /// and format of the table.
        /// </summary>
        [Input("autodetect", required: true)]
        public Input<bool> Autodetect { get; set; } = null!;

        /// <summary>
        /// The compression type of the data source.
        /// Valid values are "NONE" or "GZIP".
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        /// <summary>
        /// Additional properties to set if
        /// `source_format` is set to "CSV". Structure is documented below.
        /// </summary>
        [Input("csvOptions")]
        public Input<Inputs.TableExternalDataConfigurationCsvOptionsArgs>? CsvOptions { get; set; }

        /// <summary>
        /// Additional options if
        /// `source_format` is set to "GOOGLE_SHEETS". Structure is
        /// documented below.
        /// </summary>
        [Input("googleSheetsOptions")]
        public Input<Inputs.TableExternalDataConfigurationGoogleSheetsOptionsArgs>? GoogleSheetsOptions { get; set; }

        /// <summary>
        /// Indicates if BigQuery should
        /// allow extra values that are not represented in the table schema.
        /// If true, the extra values are ignored. If false, records with
        /// extra columns are treated as bad records, and if there are too
        /// many bad records, an invalid error is returned in the job result.
        /// The default value is false.
        /// </summary>
        [Input("ignoreUnknownValues")]
        public Input<bool>? IgnoreUnknownValues { get; set; }

        /// <summary>
        /// The maximum number of bad records that
        /// BigQuery can ignore when reading data.
        /// </summary>
        [Input("maxBadRecords")]
        public Input<int>? MaxBadRecords { get; set; }

        /// <summary>
        /// The data format. Supported values are:
        /// "CSV", "GOOGLE_SHEETS", "NEWLINE_DELIMITED_JSON", "AVRO", "PARQUET",
        /// and "DATSTORE_BACKUP". To use "GOOGLE_SHEETS"
        /// the `scopes` must include
        /// "https://www.googleapis.com/auth/drive.readonly".
        /// </summary>
        [Input("sourceFormat", required: true)]
        public Input<string> SourceFormat { get; set; } = null!;

        [Input("sourceUris", required: true)]
        private InputList<string>? _sourceUris;

        /// <summary>
        /// A list of the fully-qualified URIs that point to
        /// your data in Google Cloud.
        /// </summary>
        public InputList<string> SourceUris
        {
            get => _sourceUris ?? (_sourceUris = new InputList<string>());
            set => _sourceUris = value;
        }

        public TableExternalDataConfigurationArgs()
        {
        }
    }
}
